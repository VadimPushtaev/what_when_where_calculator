package application

import "fmt"

type GamesAnalyzer struct {
	games []*Game
}

func NewGamesAnalyzer(games []*Game) *GamesAnalyzer {
	return &GamesAnalyzer{
		games: games,
	}
}

func (analyzer *GamesAnalyzer) Analyze() {
	nameCounter := make(map[string]int)

	for _, game := range analyzer.games {
		if game == nil {
			fmt.Println("game is nil")
		}
		for _, turn := range game.history {
			sector := game.table.sectors[turn.playedSector]
			cnt, exists := nameCounter[sector.name]
			if exists {
				nameCounter[sector.name] = cnt + 1
			} else {
				nameCounter[sector.name] = 1
			}
		}
	}

	for _, sector := range analyzer.games[0].table.sectors {
		fmt.Printf("%s: %f\n", sector.name, float32(nameCounter[sector.name])/float32(len(analyzer.games)))
	}
}
