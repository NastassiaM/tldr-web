package tldr

import "net/http"

// Route stores info about particular route for http habdlers:
// route name , http method, pattern string and handler function.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is just an alias for slice of Route.
type Routes []Route

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		index,
	},
	Route{
		"tldrIndex",
		"GET",
		"/tldr",
		tldrIndex,
	},
	Route{
		"tldrItem",
		"GET",
		"/tldr/{tldrName}",
		tldrItem,
	},
	Route{
		"tldrCreate",
		"PUT",
		"/tldr",
		tldrCreate,
	},
}
