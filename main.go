package main

import (
	"formsubmit/database"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	// Initialize a new chi router
	r := chi.NewRouter()

	// Connect to the database
	err := database.Connect()
	forceOk(err)

	// Start the http server
	err = http.ListenAndServe(":3000", r)
	forceOk(err)
}

func forceOk(err error) {
	if err != nil {
		panic(err)
	}
}
