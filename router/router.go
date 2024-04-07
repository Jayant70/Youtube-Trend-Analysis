package router

import (
	"fmt"
	"github.com/gorilla/mux"
)

// NewRouter Get a router for handling your HTTP requests
func NewRouter() *mux.Router {

	fmt.Println("Creating a new Router")

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
