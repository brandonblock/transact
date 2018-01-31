package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Channeler struct {
	requests           chan http.Request
	transactions       chan Transaction
	blocks             chan Block
	lastHash           string
	activeBlock        Block
	activeTransactions []Transaction
	ticker             *time.Ticker
	mux                sync.Mutex
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
	go func() {
		for {
			select {
			case <-ch.ticker.C:
				fmt.Println("tick")
				ch.activeBlock.transactions = ch.activeTransactions
				if len(ch.activeBlock.transactions) > 0 {
					hasher := md5.New()
					ch.mux.Lock()
					hasher.Write([]byte(fmt.Sprintf("%v%d", ch.activeBlock.transactions, time.Now().Unix())))
					ch.activeBlock.blockHash = hex.EncodeToString(hasher.Sum(nil))
					ch.activeBlock.prevBlockHash = ch.lastHash
					ch.lastHash = ch.activeBlock.blockHash
					ch.blocks <- ch.activeBlock
					ch.activeTransactions = []Transaction{}
					ch.activeBlock = Block{}
					ch.mux.Unlock()
				}
			}
		}
	}()
	for t := range ch.transactions {
		ch.activeTransactions = append(ch.activeTransactions, t)
	}
}

func (ch *Channeler) write() {
	for block := range ch.blocks {
		//TODO: write to disk instead of printing
		log.Printf("%+v", block)
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
