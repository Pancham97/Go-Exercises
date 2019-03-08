// Fetches urls in parallel and reports the duration
// and the number of bytes received
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	channel := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, channel) // creates a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-channel) // printing all the stuff received from channel
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, channel chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		channel <- fmt.Sprint("While fetching %s: %v\n", url, err) // send to channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // returns byte count
	resp.Body.Close()                                 // closing resources to prevent leaking

	if err != nil {
		channel <- fmt.Sprintf("While reading %s: %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	channel <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
