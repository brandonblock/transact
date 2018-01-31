package main

import (
	"net/http"
	"time"
)

func main() {
	bus := &Channeler{
		requests:     make(chan http.Request),
		transactions: make(chan Transaction),
		lastWrite:    time.Now(),
	}
	go bus.transact()
	go bus.record()

	http.HandleFunc("/", bus.handler)
	http.ListenAndServe(":8080", nil)
}
