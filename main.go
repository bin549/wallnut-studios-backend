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
	log.Println("Starting application on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
