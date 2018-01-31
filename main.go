package main

import (
	"net/http"
)

type transaction struct {
	id        string
	key       string
	value     string
	timestamp int64
}

type block struct {
	prevBlockHash string
	blockHash     string
	transactions  []transaction
}

func main() {
	bus := &Channeler{
		requests: make(chan http.Request),
	}
	go bus.transact()

	http.HandleFunc("/", bus.handler)
	http.ListenAndServe(":8080", nil)
}

func record(transactions chan transaction) {
	return
}

func write(b block) {
	return
}
