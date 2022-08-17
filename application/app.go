package application

import (
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type App struct {
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
		logger: logger,
	}
}

func (app *App) Run() {
	app.logger.Sugar().Info("Run is started")

	rand.Seed(time.Now().UnixNano())

	games := PlayRandomGames(100_000, app.logger)
	analyzer := NewGamesAnalyzer(games)
	analyzer.Analyze()

	app.logger.Sugar().Info("Run is finished")
}
