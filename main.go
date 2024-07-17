package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"testproject/database"
	"testproject/grpc_server"
	"testproject/rest_server"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go database.CreateDatabase()
	go rest_server.StartRestServer()
	go grpc_server.StartGrpsServer()

	<-stop
	log.Println("Shutting down servers...")
}
