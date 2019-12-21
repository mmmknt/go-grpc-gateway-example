# go-grpc-gateway-example

## How to use
### gRPC only
#### Server
Move `sever` directory and execute `go run .`.

TODO how to confirm normal execution

#### Client
Move `client` directory and execute `go run .`.

You can see info log.

`2019/12/21 18:41:49 Echo: value:"Hello, World"`

### REST with grpc-gateway
#### Server
After running gRPC server, move `gateway` and execute `go run .`.

TODO how to confirm normal execution

#### Client
Send http request by curl:

`curl -D - -s -H 'Content-Type:application/json' -d '{"value":"JSON"}' -X POST http://localhost:8080/v1/example/echo`

You see normal response like this:

```
HTTP/1.1 200 OK
Content-Type: application/json
Grpc-Metadata-Content-Type: application/grpc
Date: Sat, 21 Dec 2019 09:46:49 GMT
Content-Length: 23

{"value":"Hello, JSON"}
```

### REST with gRPC-JSON transcoder
#### Server
After running gRPC server, move `envoy` directory and execute these command.

```
docker build -t service/echo -f ./envoy.Dockerfile .
docker run -d -p 8080:51051 service/echo
```

#### Client
Send http request as same as 'REST with grpc-gateway'.

You see normal response like this:

```
HTTP/1.1 200 OK
content-type: application/json
x-envoy-upstream-service-time: 1
grpc-status: 0
grpc-message: 
content-length: 28
date: Sat, 21 Dec 2019 14:57:57 GMT
server: envoy

{
 "value": "Hello, JSON"
}
```

### Regeneration
#### service.pb.go and service.pb.gw.go
```
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  --grpc-gateway_out=logtostderr=true:. \
  ./service/service.proto
```  

#### proto.pb
```
protoc -I. \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --include_imports \
  --include_source_info \
  --descriptor_set_out=./envoy/proto.pb \
  ./service/service.proto
```

