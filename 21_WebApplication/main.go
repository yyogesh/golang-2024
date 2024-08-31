package main

import (
	"log"
	"net/http"
)

type application struct {
	Domain string
}

func main() {
	// set up an app configuration
	app := application{
		Domain: "example.com",
	}
	// get application routes

	server := app.routes()

	// initialize database connection

	log.Println("Starting server on port 8080...")

	// start the server with the routes
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err)
	}
}
