package webshot

import "image"

type (
	Conv interface {
		HTMLToImage(srcHTML []byte, img *image.RGBA) error
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
