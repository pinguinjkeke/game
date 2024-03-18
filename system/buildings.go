package system

import (
	"game/component"
	"github.com/solarlune/ebitick"
	"github.com/yohamta/donburi/ecs"
	"math/rand"
	"time"
)

func UpdateBuildings(ecs *ecs.ECS) {
	timerEntry := component.Timer.MustFirst(ecs.World)
	timer := component.Timer.Get(timerEntry)
	buildingsEntry := component.Buildings.MustFirst(ecs.World)
	buildings := component.Buildings.Get(buildingsEntry)

	if buildings.Timer != nil && buildings.Timer.State != ebitick.StateFinished {
		return
	}

	buildings.Timer = timer.After(3*time.Second, func() {
		for _, building := range buildings.Buildings {
			for i := 0; i < 3; i++ {
				building.WindowLights[rand.Intn(len(building.WindowLights))] = rand.Intn(2) == 1
			}
		}
	})
}
