package application

import (
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type App struct {
	game   *Game
	logger *zap.Logger
}

func NewApp(configPath *string) *App {
	logger, _ := zap.NewProduction()

	config := NewConfig(configPath)
	logger.Sugar().Infow(
		"Configuration is set",
		"config", config,
	)

	return &App{
		game:   NewGame(logger),
		logger: logger,
	}
}

func (app *App) Run() {
	app.logger.Sugar().Info("Run is started")

	rand.Seed(time.Now().UnixNano())
	app.game.AllPossibilities()

	app.logger.Sugar().Info("Run is finished")
}
