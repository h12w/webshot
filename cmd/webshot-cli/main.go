package main

import (
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"h12.io/webshot"
)

func main() {
	src, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	img := new(image.RGBA)
	if err := webshot.NewConv().HTMLToImage(src, img); err != nil {
		log.Fatal(err)
	}
	enc := png.Encoder{CompressionLevel: png.DefaultCompression}
	if err := enc.Encode(os.Stdout, img); err != nil {
		log.Fatal(err)
	}
}
