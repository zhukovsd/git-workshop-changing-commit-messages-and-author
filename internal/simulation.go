package internal

import (
	"fmt"
	"time"
)

type Simulation struct {
	worldMap Map
}

func NewSimulation(worldMap Map) *Simulation {
	return &Simulation{
		worldMap: worldMap,
	}
}

func (s Simulation) Start() {
	for _, ia := range initActions() {
		ia.Perform(&s.worldMap)
	}
	Render(s.worldMap)

	turnAction := NewTurnAction()

	for !isOver(s.worldMap) {
		turnAction.Perform(&s.worldMap)
		Render(s.worldMap)
		time.Sleep(time.Second)
	}

	fmt.Println("Simulation is over")
}

func initActions() []Action {
	initActions := []Action{}

	initActions = append(initActions, NewSpawnAction[Herbivore](3))
	initActions = append(initActions, NewSpawnAction[Predator](2))
	initActions = append(initActions, NewSpawnAction[Grass](7))
	initActions = append(initActions, NewSpawnAction[Rock](5))
	initActions = append(initActions, NewSpawnAction[Tree](6))

	return initActions
}

func isOver(m Map) bool {
	if len(FindByType[*Herbivore](m)) == 0 {
		return true
	}
	if len(FindByType[Grass](m)) == 0 {
		return true
	}
	return false
}
