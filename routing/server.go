package routing

import (
	"net/http"
	"time"
)

// ServeHttp registers a listener which serves requests with the global router. This should be called only once.
func ServeHttp(addr string) error {
	srv := &http.Server{
		Handler:      defaultRouter,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv.ListenAndServe()
}
