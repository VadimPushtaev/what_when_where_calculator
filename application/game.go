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

type Table struct {
	sectors []Sector
}

type Game struct {
	teamA  int
	teamB  int
	goal   int
	table  Table
	logger *zap.Logger
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

func NewGame(logger *zap.Logger) *Game {
	game := &Game{
		teamA: 0,
		teamB: 0,
		goal:  2,
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

func (game *Game) AllPossibilities() {
	AllPossibilitiesRecursive(game, 1)
}

func AllPossibilitiesRecursive(game *Game, p float32) {
	if game.teamA >= game.goal || game.teamB >= game.goal {
		game.logger.Sugar().Infow(
			"Game is finished",
			"teamA", game.teamA,
			"teamB", game.teamB,
			"probability", p,
			"opened", game.OpenedSectorsNames(),
		)
		return
	}

	for i := 0; i < len(game.table.sectors); i++ {
		j := i
		for game.table.sectors[j].opened == true {
			j++
			if j >= len(game.table.sectors) {
				j = 0
			}
		}

		aGame := game.Copy()
		aGame.table.sectors[j].opened = true
		aGame.table.sectors[j].won = true
		aGame.teamA++
		AllPossibilitiesRecursive(aGame, p*game.table.sectors[j].winRate/float32(len(game.table.sectors)))

		bGame := game.Copy()
		bGame.table.sectors[j].opened = true
		aGame.table.sectors[j].won = false
		bGame.teamB++
		AllPossibilitiesRecursive(bGame, p*(1-game.table.sectors[j].winRate/float32(len(game.table.sectors))))
	}
}
