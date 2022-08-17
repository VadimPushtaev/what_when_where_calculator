package application

import (
	"math/rand"
)

type TurnPlan struct {
	selectedSector int
	loseChange     float32
}

type GamePlan struct {
	turnsPassed int
	turnPlans   []TurnPlan
}

func RandomSelectorPlan(size int) *GamePlan {
	turnPlans := make([]TurnPlan, size)
	for i := 0; i < size; i++ {
		turnPlans[i] = TurnPlan{
			selectedSector: rand.Intn(size),
			loseChange:     rand.Float32(),
		}
	}

	gamePlan := &GamePlan{
		turnsPassed: 0,
		turnPlans:   turnPlans,
	}

	return gamePlan
}

func (plan *GamePlan) Yield() TurnPlan {
	pointer := plan.turnsPassed
	plan.turnsPassed++

	return plan.turnPlans[pointer]
}
