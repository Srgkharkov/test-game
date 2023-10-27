package apiserver

import (
	"fmt"
	"github.com/Srgkharkov/test-game/internal/game"
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

var ga *game.Game

type Routes []Route

func NewRouter(g *game.Game) *mux.Router {
	ga = g
	router := mux.NewRouter().StrictSlash(true)
	//router.han
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
		strings.ToUpper("Post"),
		"/getresult",
		GetResult,
	},
}
