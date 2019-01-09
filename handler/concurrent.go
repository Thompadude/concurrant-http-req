package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Concurrent(c echo.Context) error {
	urlCount := c.Get("url_count").(int)
	urls := make([]string, urlCount)

	for i, arg := range os.Args[1:] {
		log.Debugf("Extracting URL: %s", arg)
		urls[i] = arg
	}

	start := time.Now()

	// Create the channel we will use in the doRequest function.
	ch := make(chan []byte)
	for _, url := range urls {
		go doRequest(url, ch)
	}

	// Read from the channel the same amount of times as the URL array is long.
	// We can do something with the response here if we wish.
	for range urls {
		<-ch
	}

	log.Infof("Total time elapsed for %d requests: %.2fs", len(urls), time.Since(start).Seconds())
	return nil
}

func doRequest(url string, ch chan<- []byte) { // <-- The channel can only be used to receive []byte.
	start := time.Now()
	log.Infof("Sending GET request to: %s", url)
	resp, _ := http.Get(url)
	seconds := time.Since(start).Seconds()

	bytes, _ := ioutil.ReadAll(resp.Body)
	log.Infof("Response from URL: %s ::: STATUS CODE: %d ::: TIME ELAPSED: %.2f", url, resp.StatusCode, seconds)

	// Use the channel for storing the response body.
	ch <- bytes
}
