package main

import (
	"fmt"
	"log"
	"net/http"
)

type Channeler struct {
	requests chan http.Request
}

func (ch *Channeler) handler(w http.ResponseWriter, r *http.Request) {
	ch.requests <- *r
}

func (ch *Channeler) transact() {
	for req := range ch.requests {
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

func extract(r http.Request, lookup string) (value string, err error) {
	values, ok := r.URL.Query()[lookup]
	if !ok || len(values) < 1 {
		log.Println("Url Param 'id' is missing")
		return value, fmt.Errorf("Url param %s is missing", lookup)
	}
	value = values[0]
	return
}
