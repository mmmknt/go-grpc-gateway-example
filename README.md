# go-grpc-gateway-example

## Prerequisites
TODO

## Generate codes from service.proto
Move repository root and execute command.

```
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  --grpc-gateway_out=logtostderr=true:. \
  ./service/service.proto
```  

`service.pb.go` and `service.pb.gq.go` are generated into service directory.

## Execute gRPC server
Move `sever` directory and execute command.

`go run ./`

TODO how to confirm normal execution

## Send gRPC request
Move `client` directory and execute command.

`go run ./`

You can see info log.

`2019/12/21 18:41:49 Echo: value:"Hello, World"`

## Execute reverse proxy server
Move `gateway` directory and execute command.

`go run ./`

TODO how to confirm normal execution

## Send REST request

`curl -D - -s -H 'Content-Type:application/json' -d '{"value":"JSON"}' -X POST http://localhost:8080/v1/example/echo`

You can see normal http response.

```
HTTP/1.1 200 OK
Content-Type: application/json
Grpc-Metadata-Content-Type: application/grpc
Date: Sat, 21 Dec 2019 09:46:49 GMT
Content-Length: 23

{"value":"Hello, JSON"}
```
