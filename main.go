package main

import (
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	src, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	img, err := htmlToImage(src)
	if err != nil {
		log.Fatal(err)
	}
	enc := png.Encoder{CompressionLevel: png.DefaultCompression}
	if err := enc.Encode(os.Stdout, img); err != nil {
		log.Fatal(err)
	}
}
