package application

type Table struct {
	sectors []Sector
}

func (table *Table) SelectFirstNotOpenedSector(selected int) int {
	for i := selected; ; i = (i + 1) % len(table.sectors) {
		if !table.sectors[i].opened {
			return i
		}
	}
}
