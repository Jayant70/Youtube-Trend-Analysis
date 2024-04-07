package router

import (
	"net/http"
	"youtube/constants"
	"youtube/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetVideosData",
		constants.GET,
		constants.GetVideosDataURI,
		handler.GetVideosHandler,
	},
}
