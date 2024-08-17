package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// set application config

	var app application

	// read from command line

	// initialize database connection

	app.Domain = "localhost"
	log.Println("Starting application on port", port)

	// http.HandleFunc("/", Hello)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
