package main

import (
	"net/http"
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
		"serverTest",
		"GET",
		"/",
		serverTest,
	},
	Route{
		"mainRepoIncomingHook",
		"POST",
		"/",
		mainRepoIncomingHook,
	},
	Route{
		"dependencyRepoIncomingHook",
		"POST",
		"/dep",
		dependencyRepoIncomingHook,
	},
	Route{
		"flushDatabase",
		"GET",
		"/flush",
		flushDatabase,
	},
}
