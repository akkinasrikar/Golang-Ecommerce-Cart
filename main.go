package main

import (
	golog "log"

	"github.com/akkinasrikar/ecommerce-cart/server"
	database "github.com/akkinasrikar/ecommerce-cart/database"
)

func main() {
	server, err := server.Init()
	if err != nil {
		golog.Fatal("Error initializing server", err)
	}
	golog.Println("Initialized server")
	
	database.ConnectDataBase()
	golog.Println("Connected to database")
	
	err = server.Start()
	if err != nil {
		golog.Fatal("Error starting server", err)
	}
	golog.Println("Started server")
}
