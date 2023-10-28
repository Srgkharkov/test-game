package main

import (
	"github.com/Srgkharkov/test-game/internal/apiserver"
	"github.com/Srgkharkov/test-game/internal/game"
	"github.com/Srgkharkov/test-game/internal/metrics"
	"log"
)

// Main function responsible for initializing and starting modules.
// The 'game' package contains the core logic of the service.
// The 'metrics' package includes the Prometheus client, used to initialize counters and histograms.
// The 'APIServer' package comprises the HTTP server and endpoint handlers.
func main() {
	game := game.NewGame()
	metrics := metrics.NewMetrics()
	metrics.Run()
	APIServer := apiserver.NewAPIServer(game, metrics)
	if err := APIServer.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
