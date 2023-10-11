# Honey Badger Go client

Honey Badger is simple and fast key/value database server build on top of [BadgerDB](https://github.com/dgraph-io/badger). It uses [gRPC](https://grpc.io/) as transport protocol.

## Getting Started
### Installation
```
go get github.com/honey-badger-io/go-client
```

### Usage
```go
	client, err := NewClient("127.0.0.1:18950")
	if err != nil {
		panic(err)
	}

    // Set data
    data := make([]byte, 1)
    err := client.Data(context.TODO(), "my-database").Set("some-key", data)
    if err != nil {
        panic(err)
	}

    // Get data
    data, err := client.Data(context.TODO(), "my-database").Get("some-key")
    if err != nil {
        panic(err)
	}
    fmt.Printf("%v\n", data)
```
