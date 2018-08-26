package app

import (
	"github.com/julienschmidt/httprouter"
)

const APIVersionPath = "/api/v1"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		route("/"),
		Index,
	},
	Route{
		"Health",
		"GET",
		route("/health"),
		Health,
	},
}

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle

		handle = route.HandlerFunc
		handle = Logger(handle, route.Name)

		router.Handle(route.Method, route.Path, handle)
	}

	return router
}

func route(p string) string {
	return APIVersionPath + p
}
