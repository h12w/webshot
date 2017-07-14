package main

import (
	"errors"
	"image"
	"runtime"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sourcegraph/go-webkit2/webkit2"
)

func htmlToImage(srcHTML []byte) (*image.RGBA, error) {
	var img image.RGBA
	var err error

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	gtk.Init(nil)

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

	if _, err := glib.IdleAdd(func() bool {
		webView.LoadHTML(string(srcHTML), "/src_html")
		return false
	}); err != nil {
		return nil, err
	}

	gtk.Main()

	if err == nil && img.Pix == nil {
		err = errors.New("fail to generate an image")
	}
	return &img, err
}
