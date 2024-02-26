package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"wallnut-studios.com/rest-api/internal/repository"
	"wallnut-studios.com/rest-api/internal/repository/dbrepo"
)

const port = 5000

type application struct {
	DSN    string
	DB     repository.DatabaseRepo
	Domain string
}

func main() {
	var app application
	flag.StringVar(&app.DSN, "dsn", "host=124.71.65.161 port=5432 user=postgres password=__2018bb dbname=wallnut-studios sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain")
	flag.Parse()
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	log.Println("Starting application on port", port)
	log.Println("Starting application on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
