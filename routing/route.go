package routing

import "github.com/ynori7/lilypad/handler"

// Route represents a route to be handled by the router
type Route struct {
	Path    string
	Method  string
	Host    string
	Scheme  string
	Handler handler.Handler
}

// StaticContentRoute represents a static content (e.g. images) route to be handled by the router
type StaticContentRoute struct {
	//The path prefix in the route, e.g. "/static/" will serve anything under /static/*
	PathPrefix string
	//The path to the directory to be served
	Directory string
}
