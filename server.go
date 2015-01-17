package main

import (
	"io"
	"net/http"
)


func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, r.RemoteAddr)
	})
	http.ListenAndServe(":8000", nil)
	http.ListenAndServe("[::]8000", nil)
}
