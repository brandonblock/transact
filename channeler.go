package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Channeler struct {
	requests           chan http.Request
	transactions       chan Transaction
	lastHash           string
	lastWrite          time.Time
	activeBlock        Block
	activeTransactions []Transaction
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

	for t := range ch.transactions {
		log.Printf("%v", t)
		ch.activeTransactions = append(ch.activeTransactions, t)
		//TODO: check lastWrite time constantly
		if ch.lastWrite.Before(time.Now().Add(-30 * time.Second)) {
			ch.activeBlock.transactions = ch.activeTransactions
			ch.activeBlock.blockHash = "hash" //TODO: generate hash of the block struct
			ch.activeBlock.prevBlockHash = ch.lastHash
			ch.lastHash = ch.activeBlock.blockHash
			log.Printf("%+v", ch.activeBlock)
			ch.activeTransactions = []Transaction{}
			ch.activeBlock = Block{}
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
