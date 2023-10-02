proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative honey_badger.proto

test:
	go build .
	go test . -v
