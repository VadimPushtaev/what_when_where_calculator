package application

import (
	"fmt"
	"strconv"
	"strings"
)

type SectorSetup struct {
	winRate float32
	name    string
}

type Sector struct {
	winRate float32
	name    string
	opened  bool
	won     bool
	table   *Table
}

func (sectorSetup *SectorSetup) SetValue(s string) error {
	if s == "" {
		return fmt.Errorf("field value can't be empty")
	}
	if strings.Contains(s, ":") {
		split := strings.Split(s, ":")
		if len(split) != 2 {
			return fmt.Errorf("sector setup value should be in format `name:win_rate`")
		}
		winRate, err := strconv.ParseFloat(split[1], 32)
		if err != nil {
			return fmt.Errorf("win rate should be float")
		}
		sectorSetup.name = split[0]
		sectorSetup.winRate = float32(winRate)
	} else {
		sectorSetup.name = s
		sectorSetup.winRate = 0.5
	}

	return nil
}
