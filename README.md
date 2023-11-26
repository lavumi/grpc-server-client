# gRPC Server and Client
Server made by gRPC 

## Installation

Install [protoc](http://google.github.io/proto-lens/installing-protoc.html)


## Project Layout

Refer to [golang-standards/project-layout](https://github.com/golang-standards/project-layout) 



## Run protobuf Compiler
```bash
protoc --proto_path=api/proto --go_out=pkg/api --go_opt=paths=import --go-grpc_out=pkg/api --go-grpc_opt=paths=import api/proto/types/*.proto api/proto/*.proto
```
```bash
protoc --proto_path=api/proto --js_out=import_style=commonjs:webpage/api api/proto/types/*.proto api/proto/*.proto
```

## Run gRPC Server
```bash
go run cmd/server/main.go
```

## Run gRPC Client
```bash
go run cmd/client/main.go
```


## Build 
todo improve
```bash
env GOOS=linux go build -o grpc-game-server cmd/server/main.go
docker build -t grpc-game-server:0.0.2 .
docker save -o grpc-game-server.tar grpc-game-server
rm grpc-game-server
```
