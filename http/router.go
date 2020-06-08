package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

var defaultRouter = mux.NewRouter()

// RegisterRoutes takes a slice of routes and registers them as router HandleFuncs
func RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		r := defaultRouter.HandleFunc(route.Path, getHandlerWrapper(route.Handler))

		if route.Host != "" {
			r.Host(route.Host)
		}

		if route.Scheme != "" {
			r.Schemes(route.Scheme)
		}

		if route.Method != "" {
			r.Methods(route.Method)
		}
	}
}

// RegisterStatiContentRoutes takes a slice of static content routes and registers them as router HandleFuncs
func RegisterStatiContentRoutes(routes ...StaticContentRoute) {
	for _, route := range routes {
		defaultRouter.PathPrefix(route.PathPrefix).Handler(
			http.StripPrefix(
				route.PathPrefix,
				http.FileServer(http.Dir(route.Directory)),
			),
		)
	}
}

func getHandlerWrapper(h Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h(r)

		if resp.Status == http.StatusMovedPermanently || resp.Status == http.StatusFound {
			http.Redirect(w, r, resp.RedirectURL, resp.Status)
			return
		}

		for header, value := range resp.Headers {
			w.Header().Set(header, value)
		}
		w.WriteHeader(resp.Status)
		_, _ = w.Write(resp.Body)
	}
}
