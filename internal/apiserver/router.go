package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AddConfig",
		strings.ToUpper("Post"),
		"/addconfig",
		AddConfig,
	},

	Route{
		"GetMetrics",
		strings.ToUpper("Get"),
		"/metrics",
		GetMetrics,
	},

	Route{
		"GetResult",
		strings.ToUpper("Get"),
		"/getresult",
		GetResult,
	},
}
