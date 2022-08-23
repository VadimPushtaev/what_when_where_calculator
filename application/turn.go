package application

import (
	"fmt"
	"strconv"
	"strings"
)

type GameTurn struct {
	playedSector int
	won          bool
}

func (turn *GameTurn) SetValue(s string) error {
	if s == "" {
		return fmt.Errorf("field value can't be empty")
	}
	split := strings.Split(s, ":")
	if len(split) != 2 {
		return fmt.Errorf("game turn value should be in format `number:won`")
	}
	playedSector, err := strconv.Atoi(split[0])
	if err != nil {
		return fmt.Errorf("played sector of game turn should be integer")
	}
	won, err := strconv.ParseBool(split[1])
	if err != nil {
		return fmt.Errorf("won of game turn should be boolean")
	}
	turn.playedSector = playedSector
	turn.won = won

	return nil
}
