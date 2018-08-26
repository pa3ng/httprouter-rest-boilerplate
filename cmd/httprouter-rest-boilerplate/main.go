package main

import (
	"net/http"

	"github.com/pa3ng/httprouter-boilerplate/internal/app"
)

func main() {
	router := app.NewRouter()

	err := http.ListenAndServe(":8080", http.Handler(router))
	if err != nil {
		panic(err)
	}
}
