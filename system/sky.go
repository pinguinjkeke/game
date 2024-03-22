package system

import (
	"game/component"
	"github.com/yohamta/donburi/ecs"
)

func UpdateSky(ecs *ecs.ECS) {
	skyEntry := component.Sky.MustFirst(ecs.World)
	sky := component.Sky.Get(skyEntry)

	if sky.X0 > 0 {
		sky.X0 = -sky.LayerWidth
	}

	if sky.X1 > 0 {
		sky.X1 = -sky.LayerWidth
	}

	sky.X0 += 0.05
	sky.X1 += 0.07
}
