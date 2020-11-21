package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	http.Handle("/", loggingMiddleware(http.HandlerFunc(handler)))
	port:=":8080"
	fmt.Printf("start server, listen on %s\n", port)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Infof("uri: %s", req.RequestURI)
		next.ServeHTTP(w, req)
	})
}
