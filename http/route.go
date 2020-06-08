package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents a route to be handled by the router
type Route struct {
	Path    string
	Method  string
	Host    string
	Scheme  string
	Handler Handler
}

// StaticContentRoute represents a static content (e.g. images) route to be handled by the router
type StaticContentRoute struct {
	//The path prefix in the route, e.g. "/static/" will serve anything under /static/*
	PathPrefix string
	//The path to the directory to be served
	Directory string
}

// Vars returns the route parameters
func Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

// GetVar returns a specific route parameter
func GetVar(r Request, name string) string {
	v := mux.Vars(r)
	return v[name]
}
