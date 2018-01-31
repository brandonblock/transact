package main

import (
	"net/http"
)

func main() {
	bus := &Channeler{
		requests:     make(chan http.Request),
		transactions: make(chan Transaction),
	}
	go bus.transact()
	go bus.record()

	http.HandleFunc("/", bus.handler)
	http.ListenAndServe(":8080", nil)
}
