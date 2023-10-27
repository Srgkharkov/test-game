package app

import (
	"github.com/Srgkharkov/test-game/internal/apiserver"
	"github.com/Srgkharkov/test-game/internal/game"
	"github.com/Srgkharkov/test-game/internal/metrics"
)

type App struct {
	Game      *game.Game
	APIServer *apiserver.APIServer
	Metrics   *metrics.Metrics
}

func NewApp() *App {
	game := game.NewGame()
	return &App{
		Game:      game,
		APIServer: apiserver.NewAPIServer(game),
	}
}
