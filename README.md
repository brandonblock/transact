The overall goal is to create a simple application that can receive `transactions`, append them to an in memory `block` 
and after a configured time limit commit the `block` containing a set of `transactions` to disk. 

Golang is the required language for this project. 

The project must compile and come with a set of run instructions. 

We expect that you do this independently and without using existing code. 

The mini project consists of three core parts
- receiving `transactions` through a http api. 
- storing `transactions` in memory for a configured amount of time. 
- appending a set of `transactions` ( or a `block` ) to disk when a timer has expired. 

### Project overview

Write a small http api that can receive a (key,value) pair as a request param.  

Example : "http://localhost:8080/tx?key=hello&value=world"

From the request create a `transaction` object 

A `transaction` object should consist of at least the following properties

```json
    { 
        "id" : "hash-value",
        "key": "hello", 
        "value": "world", 
        "timestamp" : "epoch-time-stamp" 
    }
```

 Append this `transaction` to a `block` object, which should contain at least the following properties

```json
    {
        "prev-block-hash": "hash-value",
        "block-hash": "hash-value",
        "transactions": [
            { "id" : "hash-value", "key": "hello", "value": "world", "timestamp" : "epoch-time-stamp" },
            { "id" : "hash-value", "key": "hi", "value": "welcome", "timestamp" : "epoch-time-stamp" }
        ]
    }
```

`Blocks` should be stored in memory for a pre configured amount of time. The application should
- receive `transaction` through the http endpoint
- either create or append `transactions` to a `block` 
- at some time interval, write a `block` to a on disk file. 

Additonally, new `blocks` should be appended to the file 

Each appended `block` should reference the previous block added (excluding the first block). 

A `block` object should be used to store incoming transactions until a timer has expired. 

When the timer has expired 
- A new `block` should be created 
- The `block` currently in memory should be writen to disk

A simple file on disk should be used to append new `blocks` as they are created. 

The above examples use json as the file/data format. However, you are welcome to use any human readable data type. 

### Bonus (not required)
- Max amount of transactions in block
    - Write to disk if max transactions is reached before timer expires and rest timer.
- Leveled logging 
- Minimal amount of dependencies 
- Dependency management 
- Modular code design
- Easy to configure
- Memory and CPU metrics 
- Clear documentation 
- Clean vcs history 
- Add search ability for key that returns a transaction
- Peer to Peer
    - simply forward transactions to other peers and allow them to commit