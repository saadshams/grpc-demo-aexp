//
//  client.go
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
	"time"
)

func main() {
	// create server connection
	connection, error := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if error != nil {
		log.Fatalf("connection errors: %v", error)
	}
	defer connection.Close()

	// create client
	client := aexp.NewAEXPClient(connection)

	// send a request and get response
	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, error := client.SendRequest(context, &aexp.Request{Name: "AEXP"})
	if error != nil {
		log.Fatalf("Failed to send request: %v", error)
	}

	log.Printf("gRPC Response: %s", response.Message)
}