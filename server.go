package main

import (
	"io"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		regex, _ := regexp.Compile(":[0-9]+$")
		addr := regex.ReplaceAllString(r.RemoteAddr, "")
		addr = strings.Replace(addr, "[", "", -1)
		addr = strings.Replace(addr, "]", "", -1)
		io.WriteString(w, addr)
	})

	http.ListenAndServe("[::]:8000", nil)
}
