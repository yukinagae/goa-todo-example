# goa-todo-example

goa todo example project

## Requirements

- go: 1.12.9

## Setup

```bash
go get -u goa.design/goa/v3
go get -u goa.design/goa/v3/...
go get -u github.com/golang/protobuf/protoc-gen-go # for gRPC
```

## Generate files

```bash
goa gen todo/design
goa example todo/design
```

## Start server

### On local

Build

```bash
go build ./cmd/todo && go build ./cmd/todo-cli
```

Run server

```bash
./todo
```

### On docker

```bash
docker-compose up
```

(Build again if Dockerfile is being changed)

```bash
docker-compose build
```

## Play with it

HTTP

```bash
./todo-cli --url="http://localhost:8080" todo get --id todo1
```

gRPC

```bash
./todo-cli --url="grpc://localhost:8888" todo get --message '{"id": "todo1"}'
```

## References

- [goa - Getting Started](https://goa.design/learn/getting-started/)
