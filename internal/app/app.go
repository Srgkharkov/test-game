package app

import (
	"github.com/Srgkharkov/test-game/internal/apiserver"
	"github.com/Srgkharkov/test-game/internal/game"
)

type App struct {
	Game *game.Game

	APIServer *apiserver.APIServer
}

func NewApp() *App {
	return &App{
		Game:      game.NewGame(),
		APIServer: apiserver.NewAPIServer(),
	}
}
