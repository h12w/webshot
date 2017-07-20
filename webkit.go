package webshot

import (
	"errors"
	"image"
	"runtime"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sourcegraph/go-webkit2/webkit2"
)

func init() {
	gtk.Init(nil)
}

type Conv struct {
	webView    *webkit2.WebView
	resultChan chan result
}

func NewConv() *Conv {
	resultChan := make(chan result)
	conv := &Conv{
		resultChan: resultChan,
		webView:    webkit2.NewWebView(),
	}
	go conv.convertingHTMLToImage()
	return conv
}

func (c *Conv) HTMLToImage(srcHTML []byte, img *image.RGBA) error {
	c.webView.LoadHTML(string(srcHTML), "/src_html")
	result := <-c.resultChan
	if result.err != nil {
		return result.err
	}
	*img = *result.img
	return nil
}

func HTMLToImage(srcHTML []byte) (*image.RGBA, error) {
	var img image.RGBA
	var err error

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	webView := webkit2.NewWebView()
	defer webView.Destroy()

	webView.Connect("load-failed", func() {
		err = errors.New("Load failed")
	})

	webView.Connect("load-changed", func(_ *glib.Object, i int) {
		loadEvent := webkit2.LoadEvent(i)
		if loadEvent == webkit2.LoadFinished {
			webView.GetSnapshot(func(result *image.RGBA, e error) {
				err = e
				if result != nil {
					img = *result
				}
				gtk.MainQuit()
			})
		}

	})

	webView.LoadHTML(string(srcHTML), "/src_html")
	gtk.Main()

	if err == nil && img.Pix == nil {
		err = errors.New("fail to generate an image")
	}
	return &img, err
}

type job struct {
	srcHTML    []byte
	resultChan chan result
}
type result struct {
	img *image.RGBA
	err error
}

func (c *Conv) convertingHTMLToImage() {
	webView, resultChan := c.webView, c.resultChan

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	webView.Connect("load-failed", func() {
		resultChan <- result{err: errors.New("Load failed")}
	})

	webView.Connect("load-changed", func(_ *glib.Object, i int) {
		loadEvent := webkit2.LoadEvent(i)
		if loadEvent == webkit2.LoadFinished {
			webView.GetSnapshot(func(img *image.RGBA, err error) {
				if err != nil {
					resultChan <- result{err: err}
					return
				}
				if img == nil || img.Pix == nil {
					resultChan <- result{err: errors.New("result image is nil")}
				}
				resultChan <- result{img: img}
			})
		}

	})

	gtk.Main()
}

func (c *Conv) Close() error {
	c.webView.Destroy()
	return nil
}
