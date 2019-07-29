## [gRPC](https://en.wikipedia.org/wiki/GRPC) [Go](https://en.wikipedia.org/wiki/Go_(programming_language)) Demo: Hello AEXP
This demo illustrates techniques for calling remote procedures in a client server application.

## Screenshot
![PureMVC Java Demo: Employee Admin Microservice](https://raw.githubusercontent.com/saadshams/grpc-demo-aexp/master/assets/screenshot.png)

## Status
Development - [Version 0.1]()

## Installation
```
go get https://github.com/saadshams/grpc-demo-aexp
cd $GOPATH/src/github.com/saadshams/grpc-demo-aexp
protoc -I aexp aexp/aexp.proto --go_out=plugins=grpc:aexp

go build server/main.go
go build client/main.go
go run server/main.go
go run client/main.go
```

## Platforms / Technologies
* [Go](https://en.wikipedia.org/wiki/Go_(programming_language))
* [gRPC](https://en.wikipedia.org/wiki/GRPC)
* [Mongodb](https://en.wikipedia.org/wiki/MongoDB)

## License
* Demo by [Saad Shams](https://www.linkedin.com/in/muizz)