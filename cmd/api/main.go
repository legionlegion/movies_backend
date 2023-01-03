package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
	DSN    string // Data source name, might be needed elsewhere asides from in connection string
	DB     repository.DatabaseRepo
}

func main() {
	// set application config
	var app application

	// read from command line
	// to connect to DB, we need connection string, to specify parameters from command line (flags)
	flag.StringVar(
		&app.DSN,
		"dsn",
		"host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTCconnect_timeout=5",
		"Postgres connection string",
	)
	flag.Parse()

	// connect to database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close() // defer runs when main returns (app closes)

	app.Domain = "example.com"

	log.Println("Starting application on port: ", port)

	http.HandleFunc("/", app.Home) // not needed after using our app.routes (chi router) as mux

	// start web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
