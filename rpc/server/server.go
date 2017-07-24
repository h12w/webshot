package server

import (
	"net"
	"net/http"
	"net/rpc"

	"h12.me/webshot/webkit"
)

func Serve(host string) error {
	rpc.Register(webkit.NewConv())
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	return http.Serve(l, nil)
}
