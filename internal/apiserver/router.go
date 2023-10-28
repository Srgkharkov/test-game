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

// The NewRouter function initializes the router, repository, and also initializes and defines handlers for the endpoints.
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

		Route{
			"GetResult",
			strings.ToUpper("Post"),
			"/getresult",
			h.GetResult,
		},
	}

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = h.Observer(handler)

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

// The APIHandler structure includes a repository for the handlers.
type APIHandler struct {
	game    *game.Game
	metrics *metrics.Metrics
}

// The NewAPIHandler function initializes the repository for the handlers.
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
