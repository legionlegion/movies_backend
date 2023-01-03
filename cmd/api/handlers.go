package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, world from %s", app.Domain);
	var payload = struct {
		Status string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status: "active",
		Message: "Go Movies up and runnning",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK);
	w.Write(out);
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies();

	if err != nil {
		fmt.Println(err)
		return
	}

	// rd, _ := time.Parse("2006-01-02", "1986-03-07")
	// rd2, _ := time.Parse("2006-01-02", "2022-12-12")

	// highlander := models.Movie {
	// 	ID: 1,
	// 	Title: "Highlander",
	// 	ReleaseDate: rd,
	// 	MPAARating: "R",
	// 	RunTime: 116,
	// 	Description: "A movie I have not watched before",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

	// avatar := models.Movie {
	// 	ID: 2,
	// 	Title: "Avatar 2",
	// 	ReleaseDate: rd2,
	// 	MPAARating: "R",
	// 	RunTime: 180,
	// 	Description: "A nice movie",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

	// movies = append(movies, highlander);
	// movies = append(movies, avatar);

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK);
	w.Write(out);
}