package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"time"

	"h12.me/webshot/rpc/server"
)

func main() {
	const xvfbPort = ":98"
	os.Setenv("DISPLAY", xvfbPort)
	errBuf := new(bytes.Buffer)
	xvfb := exec.Command("/usr/bin/Xvfb",
		xvfbPort,
		"-screen", "0",
		"2048x2048x24")
	xvfb.Stderr = errBuf
	if err := xvfb.Start(); err != nil {
		log.Fatalf("fail to start Xvfb(%v): %s", err, errBuf.String())
	}
	time.Sleep(2 * time.Second)
	if err := server.Serve(":9191"); err != nil {
		log.Fatal(err)
	}
}
