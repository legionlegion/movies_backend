package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter();

	// if application panics, log with a backtrace with necessary header
	mux.Use(middleware.Recoverer);
	mux.Use(app.enableCORS);

	mux.Get("/", app.Home);
	mux.Get("/movies", app.AllMovies);

	return mux;
}