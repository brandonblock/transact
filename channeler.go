package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Channeler struct {
	requests     chan http.Request
	transactions chan Transaction
	lastHash     string
	lastWrite    time.Time
}

type Transaction struct {
	id        string
	key       string
	value     string
	timestamp int64
}

type Block struct {
	prevBlockHash string
	blockHash     string
	transactions  []Transaction
}

func (ch *Channeler) handler(w http.ResponseWriter, r *http.Request) {
	ch.requests <- *r
}

func (ch *Channeler) transact() {
	//TODO: Error handling and logging
	for req := range ch.requests {
		id, err := extract(req, "id")
		key, err := extract(req, "key")
		value, err := extract(req, "value")
		ts, err := extract(req, "timestamp")
		timestamp, err := strconv.ParseInt(ts, 10, 64)
		if err != nil {
			panic(err)
		}
		if err != nil {
			log.Println(err.Error())
		}

		t := Transaction{
			id:        id,
			key:       key,
			value:     value,
			timestamp: timestamp,
		}
		ch.transactions <- t
	}
}

func (ch *Channeler) record() {

	ts := []Transaction{}
	for t := range ch.transactions {
		log.Printf("%v", t)
		ts = append(ts, t)
		if ch.lastWrite.Before(time.Now().Add(-30 * time.Second)) {
			//TODO: Make this block accesible outside this clause
			b := Block{}
			b.transactions = ts
			log.Printf("%v", b)
			ch.lastWrite = time.Now()
		}
	}
}

func (ch *Channeler) write(b Block) {
	return
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
