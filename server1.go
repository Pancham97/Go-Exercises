// It's a minimal echo server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls the handler function
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the path of the request URL
func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
