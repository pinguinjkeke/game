package system

import (
	"game/component"
	"github.com/yohamta/donburi/ecs"
)

func UpdateTimer(ecs *ecs.ECS) {
	timerEntry := component.Timer.MustFirst(ecs.World)
	timer := component.Timer.Get(timerEntry)

	timer.Update()
}
