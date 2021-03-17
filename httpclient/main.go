// https://tleyden.github.io/blog/2016/11/21/tuning-the-go-http-client-library-for-load-testing/?ref=hackernoon.com
package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	httpClient httpClient
}

func NewHttpClient() *client {
	return &client{
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				MaxConnsPerHost: 3,
				MaxIdleConns: 0,
				MaxIdleConnsPerHost: 3,
				IdleConnTimeout: 1 * time.Millisecond,
			},
			Timeout: 10 * time.Second,
		},
	}
}

func (c *client) Do(host string) (int, error) {
	request, err := http.NewRequest(http.MethodGet, host, nil)
	if err != nil {
		return 0, err
	}
	res, err := c.httpClient.Do(request)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	return res.StatusCode, err
}

func TestMultipleConnections() {
	n := 5
	c := NewHttpClient()
	var wg sync.WaitGroup
	for i:=0; i<n; i++{
		wg.Add(1)
		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done()
			res, err := c.Do("https://www.google.com")
			if err != nil {
				fmt.Printf("error taskId=%d err: %s\n", id, err)
				return
			}
			if res == http.StatusOK {
				fmt.Printf("successfully fetch data [taskId=%d]\n", id)
			} else {
				fmt.Printf("Error fetching data [taskId=%d] httCode:%d\n", id, res)
			}
		}(i, &wg)
	}
	wg.Wait()
	fmt.Println("finish getting results")
}

func main() {
	TestMultipleConnections()
}