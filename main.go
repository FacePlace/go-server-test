package main

import (
	"log"
	"net/http"
	"time"
)

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	toSend := []byte("<html><head></head><body>Hello</hello></html>")
	_, err := w.Write(toSend)
	if err != nil {
		log.Printf("Received an error while writing HTML: %s", err)
	}
}

func main() {
	myServer := &http.Server{
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      &myHandler{},
	}

	log.Fatal(myServer.ListenAndServe())
}
