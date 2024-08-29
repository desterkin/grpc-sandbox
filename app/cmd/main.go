package main

import (
	"log"
	"grpc_inventory/service"
)

func main() {
	log.Printf("Main!")

	service.NewService()
}

