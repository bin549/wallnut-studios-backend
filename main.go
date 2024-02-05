package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 5000

type application struct {
	Domain string
}

func main() {
	var app application
	flag.StringVar(&app.Domain, "domain", "example.com", "domain")
	flag.Parse()
	log.Println("Starting application on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
