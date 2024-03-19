package system

import (
	"game/component"
	"game/factory"
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
				window := rand.Intn(len(building.Windows))
				lights, nextLights := building.Windows[window].Lights, rand.Intn(2) == 1

				if lights == nextLights {
					continue
				}

				building.Windows[window].Lights = nextLights
				factory.RenderLights(building, window)
			}
		}
	})
}
