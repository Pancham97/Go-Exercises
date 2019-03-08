package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", req.Method, req.URL, req.Proto)

	for key, value := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", key, value)
	}

	fmt.Fprintf(w, "Host = %q\n", req.Host)
	fmt.Fprintf(w, "Remote Addr = %q\n", req.RemoteAddr)

	if err := req.ParseForm; err != nil {
		log.Print(err)
	}

	for key, value := range req.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", key, value)
	}
}

func counter(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}
