package application

type Table struct {
	sectors []Sector
}

func GetDefaultSectorSetups() []SectorSetup {
	return []SectorSetup{
		{winRate: 0.5, name: "A"},
		{winRate: 0.5, name: "B"},
		{winRate: 0.5, name: "C"},
		{winRate: 0.5, name: "D"},
		{winRate: 0.4, name: "Blitz"},
		{winRate: 0.5, name: "E"},
		{winRate: 0.5, name: "F"},
		{winRate: 0.5, name: "G"},
		{winRate: 0.2, name: "Superblitz"},
		{winRate: 0.5, name: "H"},
		{winRate: 0.5, name: "I"},
		{winRate: 0.5, name: "J"},
		{winRate: 0.5, name: "13"},
	}
}

func (table *Table) SelectFirstNotOpenedSector(selected int) int {
	for i := selected; ; i = (i + 1) % len(table.sectors) {
		if !table.sectors[i].opened {
			return i
		}
	}
}
