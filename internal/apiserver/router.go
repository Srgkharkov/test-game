package apiserver

import (
	"fmt"
	"github.com/Srgkharkov/test-game/internal/game"
	"github.com/Srgkharkov/test-game/internal/metrics"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func NewRouter(game *game.Game, metrics *metrics.Metrics) *mux.Router {
	h := NewAPIHandler(game, metrics)

	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			h.Index,
		},

		Route{
			"AddConfig",
			strings.ToUpper("Post"),
			"/addconfig",
			h.AddConfig,
		},

		//Route{
		//	"GetMetrics",
		//	strings.ToUpper("Get"),
		//	"/metrics",
		//	//h.GetMetrics,
		//	h.Index,
		//},

		Route{
			"GetResult",
			strings.ToUpper("Post"),
			"/getresult",
			h.GetResult,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	//router.han
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		//handler = h.Logger(handler, route.Name)
		handler = h.Observer(handler)
		//handler = promhttp.Handler()
		//handler = route.HandlerFunc
		//if route.UseCounter {
		//	handler = h.Logger(handler, route.Name)
		//}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	var handler http.Handler
	handler = promhttp.Handler()
	router.
		Methods(strings.ToUpper("Get")).
		Path("/metrics").
		Name("GetMetrics").
		Handler(handler)

	return router
}

type APIHandler struct {
	game    *game.Game
	metrics *metrics.Metrics
}

func NewAPIHandler(game *game.Game, metrics *metrics.Metrics) (h *APIHandler) {
	return &APIHandler{
		game:    game,
		metrics: metrics,
	}
}

//func NewRouter(g *game.Game) *mux.Router {
//	h := NewAPIHandler(g)
//	router := mux.NewRouter().StrictSlash(true)
//
//	var tmp http.HandlerFunc = h.AddConfig
//	//klh.
//	//
//	var handler http.Handler
//	//
//	handler = tmp
//	handler = Logger(handler, routes[1].Name)
//
//	rcon := Route{
//		"AddConfig",
//		strings.ToUpper("Post"),
//		"/addconfig",
//		h.AddConfig,
//	}
//	//handler = AddRepo(handler, g)
//	//h.AddConfig()
//	router.
//		Methods("POST").
//		Path("/addconfig").
//		Name("AddConfig").
//		Handler(rcon.HandlerFunc)
//	//router.han
//	//for _, route := range routes {
//	//	var handler http.Handler
//	//	handler = route.HandlerFunc
//	//	handler = Logger(handler, route.Name)
//	//
//	//	router.
//	//		Methods(route.Method).
//	//		Path(route.Pattern).
//	//		Name(route.Name).
//	//		Handler(handler)
//	//}
//
//	return router
//}

func (h *APIHandler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
