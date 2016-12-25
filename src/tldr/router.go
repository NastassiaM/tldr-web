package tldr

import "github.com/gorilla/mux"

// NewRouter creates a new gorilla/mux http handler and assigns
// precreated routes to it.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
