package main

import (
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
)

func main() {
	html, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var img image.RGBA
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9191")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	err = client.Call("Conv.HTMLToImage", html, &img)
	if err != nil {
		log.Fatal("conv error:", err)
	}

	enc := png.Encoder{CompressionLevel: png.DefaultCompression}
	if err := enc.Encode(os.Stdout, &img); err != nil {
		log.Fatal(err)
	}
}
