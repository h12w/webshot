package client

import (
	"image"
	"net/rpc"

	"h12.me/webshot"
)

type (
	Client struct {
		jobChan chan webshot.Job
	}
)

func New(addrs ...string) (*Client, error) {
	jobChan := make(chan webshot.Job)
	for _, addr := range addrs {
		rpcClient, err := rpc.DialHTTP("tcp", addr)
		if err != nil {
			return nil, err
		}
		go serve(rpcClient, jobChan)
	}
	return &Client{jobChan: jobChan}, nil
}

func (c *Client) HTMLToImage(srcHTML []byte) (*image.RGBA, error) {
	resultChan := make(chan webshot.Result)
	c.jobChan <- webshot.Job{
		SrcHTML:    srcHTML,
		ResultChan: resultChan,
	}
	result := <-resultChan
	return result.Img, result.Err
}

func serve(client *rpc.Client, jobChan chan webshot.Job) error {
	for job := range jobChan {
		var img image.RGBA
		if err := client.Call("Conv.HTMLToImage", job.SrcHTML, &img); err != nil {
			job.ResultChan <- webshot.Result{Err: err}
			continue
		}
		job.ResultChan <- webshot.Result{Img: &img}
	}
	return nil
}
