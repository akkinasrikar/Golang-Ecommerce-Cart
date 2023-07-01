package main

import (
	golog "log"

	"github.com/akkinasrikar/ecommerce-cart/server"
)

func main() {
	server, err := server.Init()
	if err != nil {
		golog.Fatal("Error initializing server", err)
	}
	golog.Println("Initialized server")
	err = server.Start()
	if err != nil {
		golog.Fatal("Error starting server", err)
	}
	golog.Println("Started server")
}
