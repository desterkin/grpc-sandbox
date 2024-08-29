package main

import (
	"grpc_inventory/service"
	"log"
	"net"

	"google.golang.org/grpc/reflection"
)

func main() {
	port := "8080"

	address := "0.0.0.0:" + port

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error starting listener on %s: %v", port, err)
	}

	log.Printf("Server started. Listening on %s", port)

	server := service.NewInventoryServer()
	reflection.Register(server)

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}
}
