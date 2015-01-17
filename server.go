package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
	http.ListenAndServe("[::]8000", nil)
}
