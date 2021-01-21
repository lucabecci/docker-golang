package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/lucabecci/docker-golang/internal/server"
)

//Simple example with channels for my docker app
func main() {
	port := "3000"

	srv, err := server.New(port)

	if err != nil {
		log.Fatal(err)
	}

	go srv.Start()

	//If the developer use the Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	//Shutdown
	srv.Close()
}
