package main

import (
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"h12.io/webshot/rpc/client"
)

func main() {
	html, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	client, err := client.New("127.0.0.1:9191")
	img, err := client.HTMLToImage(html)
	if err != nil {
		log.Fatal("conv error:", err)
	}
	enc := png.Encoder{CompressionLevel: png.DefaultCompression}
	if err := enc.Encode(os.Stdout, img); err != nil {
		log.Fatal(err)
	}
}
