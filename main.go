package main

import (
	"fmt"
	"log"
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

type CommsBus struct {
	requests chan http.Request
}

func main() {
	bus := &CommsBus{
		requests: make(chan http.Request),
	}
	go bus.transact()

	http.HandleFunc("/", bus.handler)
	http.ListenAndServe(":8080", nil)
}

func (b *CommsBus) handler(w http.ResponseWriter, r *http.Request) {
	b.requests <- *r
}

func extract(r http.Request, lookup string) (value string, err error) {
	values, ok := r.URL.Query()[lookup]
	if !ok || len(values) < 1 {
		log.Println("Url Param 'id' is missing")
		return value, fmt.Errorf("Url param %s is missing", lookup)
	}
	value = values[0]
	return
}

func (b *CommsBus) transact() {
	for req := range b.requests {
		id, err := extract(req, "id")
		key, err := extract(req, "key")
		value, err := extract(req, "value")
		timestamp, err := extract(req, "timestamp")
		if err != nil {
			log.Println(err.Error())
		}

		log.Printf("Url Param id: %s, key: %s, value: %s, timestamp: %s", id, key, value, timestamp)
	}
}

func record(transactions chan transaction) {
	return
}

func write(b block) {
	return
}
