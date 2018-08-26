package app

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// A Logger function which simply wraps the handler function around some log messages
func Logger(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params), name string) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()

		fn(w, r, param)

		log.Printf(
			"%s: %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	}
}
