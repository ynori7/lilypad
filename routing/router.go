package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

var defaultRouter = mux.NewRouter()

// RegisterRoutes takes a slice of routes and registers them as router HandleFuncs
func RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		r := defaultRouter.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			resp := route.Handler(r)

			if resp.Status == http.StatusMovedPermanently || resp.Status == http.StatusFound {
				http.Redirect(w, r, resp.RedirectUrl, resp.Status)
				return
			}

			w.WriteHeader(resp.Status)
			w.Write([]byte(resp.Body))
		})

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

// RegisterRoutes takes a slice of routes and registers them as router HandleFuncs
func RegisterStaticontentRoutes(routes ...StaticContentRoute) {
	for _, route := range routes {
		defaultRouter.PathPrefix(route.PathPrefix).Handler(
			http.StripPrefix(
				route.PathPrefix,
				http.FileServer(http.Dir(route.Directory)),
			),
		)
	}
}
