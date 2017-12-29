# Go-Blockchain

This is just a simple implementation of a blockchain in go.

## Endpoints

It provides endpoints for:

* `POST /mine` Endppoint to mine current block
* `POST /transaction` add a new transaction to the current block in the blockchain
* `GET  /chain` endpoint to get the complete chain


## Installation

Make sure you have `dep` installed. https://github.com/golang/dep

```
dep ensure
go build main.go
go run main.go
```


