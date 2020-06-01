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
