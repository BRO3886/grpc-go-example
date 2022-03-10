# Go GRPC Simple example

## Pre requisites
- protoc compilers to re-gen the proto files

## To generate code from .proto file

```bash
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
```

```bash
protoc --go_out=. --go_opt=paths=source_relative \
 --go-grpc_out=. --go-grpc_opt=paths=source_relative \
 ./addressbook/addressbook.proto
```

## Running

```bash
$ go run server.go
```

Followed by

```bash
go run client.go --option [create, add, list]
```
