package main

import (
	"log"

	"h12.me/webshot/rpc/server"
)

func main() {
	if err := server.Serve(":9191"); err != nil {
		log.Fatal(err)
	}
}
