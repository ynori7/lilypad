package http

import (
	"net/http"
	"time"
)

// ServeHttp registers a listener which serves requests with the global router. This should be called only once.
func ServeHttp(addr string, opts ...ServerOption) error {
	defaultOptions := ServerOptions{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	for _, opt := range opts {
		opt(&defaultOptions)
	}

	srv := &http.Server{
		Handler:      defaultRouter,
		Addr:         addr,
		WriteTimeout: defaultOptions.WriteTimeout,
		ReadTimeout:  defaultOptions.ReadTimeout,
	}
	return srv.ListenAndServe()
}

// ServerOption is a function that configures the HTTP server. It is used to configure the HTTP server before it is started.
type ServerOption func(*ServerOptions)

// ServerOptions are the configurable options for the HTTP server
type ServerOptions struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// WithReadTimeout sets the read timeout for the HTTP server.
func WithReadTimeout(timeout time.Duration) func(*ServerOptions) {
	return func(opts *ServerOptions) {
		opts.ReadTimeout = timeout
	}
}

// WithWriteTimeout sets the write timeout for the HTTP server.
func WithWriteTimeout(timeout time.Duration) func(*ServerOptions) {
	return func(opts *ServerOptions) {
		opts.WriteTimeout = timeout
	}
}
