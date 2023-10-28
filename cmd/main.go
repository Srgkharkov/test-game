package main

import (
	"github.com/Srgkharkov/test-game/internal/apiserver"
	"github.com/Srgkharkov/test-game/internal/game"
	"github.com/Srgkharkov/test-game/internal/metrics"
	"log"
)

func main() {
	game := game.NewGame()
	metrics := metrics.NewMetrics()
	metrics.Run()
	APIServer := apiserver.NewAPIServer(game, metrics)
	if err := APIServer.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
