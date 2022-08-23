package application

import (
	"errors"
	"go.uber.org/zap"
	"strconv"
)

type Game struct {
	teamA   int
	teamB   int
	goal    int
	table   Table
	history []*GameTurn
	logger  *zap.Logger
}

func (game *Game) OpenedSectorsNames() []string {
	var openedSectorsNames []string
	for i := 0; i < len(game.table.sectors); i++ {
		if game.table.sectors[i].opened {
			prefix := "-"
			if game.table.sectors[i].won {
				prefix = "+"
			}
			openedSectorsNames = append(openedSectorsNames, prefix+game.table.sectors[i].name)
		}
	}

	return openedSectorsNames
}

func (game *Game) AddSector(winRate float32, name string) {
	game.table.sectors = append(game.table.sectors, Sector{
		winRate: winRate,
		name:    name,
		opened:  false,
		table:   &game.table,
	})
}

func (game *Game) AddOpenedSector(winRate float32, name string) {
	game.table.sectors = append(game.table.sectors, Sector{
		winRate: winRate,
		name:    name,
		opened:  true,
		table:   &game.table,
	})
}

func NewGame(goal int, sectorSetups []SectorSetup, alreadyPlayed []GameTurn, logger *zap.Logger) (*Game, error) {
	game := &Game{
		teamA:  0,
		teamB:  0,
		goal:   goal,
		table:  Table{},
		logger: logger,
	}

	for i := 0; i < len(sectorSetups); i++ {
		game.AddSector(sectorSetups[i].winRate, sectorSetups[i].name)
	}

	for _, turn := range alreadyPlayed {
		err := game.PlayTurn(&turn)
		if err != nil {
			return nil, err
		}
	}

	if len(game.table.sectors) < goal*2-1 {
		return nil, errors.New("Not enough sectors for this goal")
	}

	return game, nil
}

func (game *Game) PlayTurn(turn *GameTurn) error {
	// Can't play non-existing sector
	if turn.playedSector >= len(game.table.sectors) {
		return errors.New(
			"Can't play sector " + strconv.Itoa(turn.playedSector) +
				" out of " + strconv.Itoa(len(game.table.sectors)),
		)
	}

	// Can't play the same sector again
	if game.table.sectors[turn.playedSector].opened {
		return errors.New("can't play the same sector again")
	}

	game.history = append(game.history, turn)
	game.table.sectors[turn.playedSector].opened = true
	game.table.sectors[turn.playedSector].won = turn.won
	if turn.won {
		game.teamA++
	} else {
		game.teamB++
	}

	return nil
}

func PlayRandomGames(config *AppConfiguration, logger *zap.Logger) ([]*Game, error) {
	var result []*Game

	for i := 0; i < config.N; i++ {
		game, err := NewGame(config.Goal, config.Sectors, config.PlayedTurns, logger)
		if err != nil {
			return nil, err
		}

		game.PlayRandom()
		result = append(result, game)
	}

	return result, nil
}

func (game *Game) PlayRandom() {
	plan := RandomSelectorPlan(len(game.table.sectors))
	game.PlayByPlan(plan)
}

func (game *Game) PlayByPlan(plan *GamePlan) {
	var p float32 = 1.0
	for {
		if game.teamA >= game.goal || game.teamB >= game.goal {
			return
		}

		turnPlan := plan.Yield()
		selected := game.table.SelectFirstNotOpenedSector(turnPlan.selectedSector)
		won := turnPlan.loseChange < game.table.sectors[selected].winRate
		chanceOfThisGameResult := game.table.sectors[selected].winRate
		if !won {
			chanceOfThisGameResult = 1 - chanceOfThisGameResult
		}
		p = p * chanceOfThisGameResult / float32(len(game.table.sectors))

		turn := &GameTurn{
			playedSector: selected,
			won:          won,
		}
		game.PlayTurn(turn)
	}
}
