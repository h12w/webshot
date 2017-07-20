package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"h12.me/webshot"
)

/*
type Conv struct{}

func (*Conv) HTMLToImage(html []byte, reply *image.RGBA) error {
	img, err := webshot.HTMLToImage(html)
	if err != nil {
		return err
	}
	*reply = *img
	return nil
}
*/

func main() {
	rpc.Register(webshot.NewConv())
	rpc.HandleHTTP()
	log.Print("listening on :9191")
	l, e := net.Listen("tcp", ":9191")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
