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
		activeBlock:  Block{},
	}
	go bus.transact()
	go bus.record()

	http.HandleFunc("/tx", bus.handler)
	http.ListenAndServe(":8080", nil)
}
