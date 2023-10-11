# Honey Badger Go client

Honey Badger is simple and fast key/value database server build on top of [BadgerDB](https://github.com/dgraph-io/badger). It uses [gRPC](https://grpc.io/) as transport protocol.

## Getting Started
### Installation
```
go get github.com/honey-badger-io/go-client
```

### Usage
```go
package main

import (
	"context"
	"fmt"

	hb "github.com/honey-badger-io/go-client"
)

func main() {
	client, err := hb.NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}

	errSet := client.Data(context.Background(), "database").Set("some-key", []byte("some data"))
	if errSet != nil {
		panic(errSet)
	}

	data, hit, errGet := client.Data(context.Background(), "database").Get("some-key")
	if errGet != nil {
		panic(errGet)
	}

	fmt.Printf("Data: %v\n", string(data))
	fmt.Printf("Hit: %v\n", hit)
}
```
### Run demo
First you must run server. Either run it from [source code](https://github.com/meeron/honey-badger) or run as docker container:
```
docker run --name honey-badger -p 18950:18950 -d meeron/honey-badger:latest
```
Then run demo program:
```sh
$ go run .
Data: some data
Hit: true
```
