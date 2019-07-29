//
//  server.go
//  gRPC Go Demo
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package main

import (
	"context"
	"github.com/saadshams/grpc-demo-aexp/aexp"
	"google.golang.org/grpc"
	"log"
	"net"
)

// implementing object
type server struct{
}

// implement rpc method
func (self *server) SendRequest(context context.Context, request *aexp.Request) (*aexp.Response, error) {
	log.Printf("Incoming request: %v", request.Name)
	return &aexp.Response{Message:"Hello " + request.Name}, nil
}

func main() {
	listener, error := net.Listen("tcp", ":50051")
	if error != nil {
		log.Fatalf("failed to start server: %v", error)
	}

	grpc := grpc.NewServer()
	aexp.RegisterAEXPServer(grpc, &server{})
	if err := grpc.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}