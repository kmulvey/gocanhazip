package main

import (
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

func main() {
	var log = logrus.New()
	log.Level = logrus.InfoLevel
	log.Formatter = &logrus.TextFormatter{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var start time.Time = time.Now()
		var regex *regexp.Regexp
		var addr string

		regex, _ = regexp.Compile(":[0-9]+$")
		addr = regex.ReplaceAllString(r.RemoteAddr, "")
		addr = strings.Replace(addr, "[", "", -1)
		addr = strings.Replace(addr, "]", "", -1)

		// send it out
		io.WriteString(w, addr)

		// time it and log it
		latency := time.Since(start)
		log.WithFields(logrus.Fields{
			"method":  r.Method,
			"request": r.RequestURI,
			"remote":  addr,
			"took":    latency,
		}).Info()
	})

	http.ListenAndServe("[::]:8000", nil)
}
