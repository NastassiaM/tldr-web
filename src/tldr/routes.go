package tldr

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TLDRIndex",
		"GET",
		"/tldr",
		TLDRIndex,
	},
	Route{
		"TLDRItem",
		"GET",
		"/tldr/{tldrName}",
		TLDRItem,
	},
	Route{
		"TLDRCreate",
		"POST",
		"/tldr",
		TLDRCreate,
	},
}
