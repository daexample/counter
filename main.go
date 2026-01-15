package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var counter uint64

func handler(w http.ResponseWriter, r *http.Request) {
	n := atomic.AddUint64(&counter, 1)
	fmt.Fprint(w, n)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
