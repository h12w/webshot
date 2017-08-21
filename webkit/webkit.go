package webkit

import (
	"errors"
	"image"
	"runtime"
	"sync"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sourcegraph/go-webkit2/webkit2"
)

var startGTKOnce sync.Once

type (
	Conv struct {
		webView    *webkit2.WebView
		resultChan chan Result
		mu         sync.Mutex
	}
	Job struct {
		SrcHTML    []byte
		ResultChan chan Result
	}
	Result struct {
		Img *image.RGBA
		Err error
	}
)

func NewConv() *Conv {
	startGTKOnce.Do(func() {
		go func() {
			runtime.LockOSThread()
			gtk.Init(nil)
			gtk.Main()
		}()
	})

	resultChan := make(chan Result)
	webViewChan := make(chan *webkit2.WebView)
	glib.IdleAdd(func() bool {
		webView := webkit2.NewWebView()

		webView.Connect("load-failed", func() {
			resultChan <- Result{Err: errors.New("load failed")}
		})

		webView.Connect("load-changed", func(_ *glib.Object, i int) {
			loadEvent := webkit2.LoadEvent(i)
			if loadEvent == webkit2.LoadFinished {
				webView.GetSnapshot(func(img *image.RGBA, err error) {
					if err != nil {
						resultChan <- Result{Err: err}
						return
					}
					if img == nil || img.Pix == nil {
						resultChan <- Result{Err: errors.New("result image is nil")}
					}
					resultChan <- Result{Img: img}
				})
			}
		})
		webViewChan <- webView
		return false
	})
	conv := &Conv{
		webView:    <-webViewChan,
		resultChan: resultChan,
	}
	return conv
}

func (c *Conv) HTMLToImage(srcHTML []byte, img *image.RGBA) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	glib.IdleAdd(func() bool {
		c.webView.LoadHTML(string(srcHTML), "/src_html")
		return false
	})
	result := <-c.resultChan
	if result.Err != nil {
		return result.Err
	}
	*img = *result.Img
	return nil

}

func (c *Conv) Close() error {
	c.webView.Destroy()
	return nil
}
