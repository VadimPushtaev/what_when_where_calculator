package application

import (
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type App struct {
	config *AppConfiguration
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
		config: config,
		logger: logger,
	}
}

func (app *App) Run() {
	app.logger.Sugar().Info("Run is started")

	rand.Seed(time.Now().UnixNano())

	games, err := PlayRandomGames(app.config, app.logger)
	if err != nil {
		app.logger.Sugar().Errorw("Error while playing games", "error", err)
	} else {
		analyzer := NewGamesAnalyzer(games)
		analyzer.Analyze()
	}

	app.logger.Sugar().Info("Run is finished")
}
