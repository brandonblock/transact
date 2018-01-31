package main

import (
	"net/http"
	"time"
)

func main() {
	ch := &Channeler{
		requests:     make(chan http.Request),
		transactions: make(chan Transaction),
		blocks:       make(chan Block),
		activeBlock:  Block{},
		ticker:       time.NewTicker(time.Second * 10),
	}
	go ch.transact()
	go ch.record()
	go ch.write()

	http.HandleFunc("/tx", ch.handler)
	http.ListenAndServe(":8080", nil)
}
