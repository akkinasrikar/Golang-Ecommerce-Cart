package main

import (
	golog "log"

	"github.com/akkinasrikar/ecommerce-cart/config"
	database "github.com/akkinasrikar/ecommerce-cart/database"
	"github.com/akkinasrikar/ecommerce-cart/server"
)

func main() {
	db := database.ConnectDataBase()
	golog.Println("Connected to database")
	config.Init()
	golog.Println("Initialized config")

	server, err := server.Init(db)
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
