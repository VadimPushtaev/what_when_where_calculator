package application

import (
	"go.uber.org/zap"
)

type Sector struct {
	winRate float32
	name    string
	opened  bool
	won     bool
	table   *Table
}

type GameTurn struct {
	selectedSector int
	won            bool
	probability    float32
}

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

func NewGame(logger *zap.Logger) *Game {
	game := &Game{
		teamA: 0,
		teamB: 0,
		goal:  6,
		table: Table{
			[]Sector{},
		},
		logger: logger,
	}

	game.AddSector(0.5, "A")
	game.AddSector(0.5, "B")
	game.AddSector(0.5, "C")
	game.AddSector(0.5, "D")
	game.AddSector(0.4, "Blitz")
	game.AddSector(0.5, "E")
	game.AddSector(0.5, "F")
	game.AddSector(0.5, "G")
	game.AddSector(0.2, "Superblitz")
	game.AddSector(0.5, "H")
	game.AddSector(0.5, "I")
	game.AddSector(0.5, "J")
	game.AddSector(0.5, "13")

	return game
}

func (game *Game) Copy() *Game {
	var newSectors []Sector
	newSectors = append(newSectors, game.table.sectors...)

	return &Game{
		teamA: game.teamA,
		teamB: game.teamB,
		goal:  game.goal,
		table: Table{
			sectors: newSectors,
		},
		logger: game.logger,
	}
}

func PlayRandomGames(n int, logger *zap.Logger) []*Game {
	var result []*Game

	for i := 0; i < n; i++ {
		game := NewGame(logger)
		game.PlayRandom()
		result = append(result, game)
	}

	return result
}

func (game *Game) PlayRandom() {
	plan := RandomSelectorPlan(len(game.table.sectors))
	game.PlayByPlan(plan)
}

func (game *Game) PlayByPlan(plan *GamePlan) {
	var p float32 = 1.0
	for {
		if game.teamA >= game.goal || game.teamB >= game.goal {
			game.logger.Sugar().Infow(
				"Game is finished",
				"teamA", game.teamA,
				"teamB", game.teamB,
				"probability", p,
				"opened", game.OpenedSectorsNames(),
			)
			return
		} else {
			game.logger.Sugar().Infow(
				"Game is not finished",
				"teamA", game.teamA,
				"teamB", game.teamB,
				"probability", p,
				"opened", game.OpenedSectorsNames(),
			)
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
			selectedSector: selected,
			won:            won,
			probability:    p,
		}
		game.history = append(game.history, turn)

		game.table.sectors[turn.selectedSector].opened = true
		game.table.sectors[turn.selectedSector].won = won
		if won {
			game.teamA++
		} else {
			game.teamB++
		}
	}
}
