package system

import (
	"game/component"
	"github.com/yohamta/donburi/ecs"
	"time"
)

var lastTime = 0.0
var accumulator = 0.0

func UpdateSpace(ecs *ecs.ECS) {
	spaceEntry := component.Space.MustFirst(ecs.World)
	space := component.Space.Get(spaceEntry)

	newTime := float64(time.Now().UnixNano()) / 1.e9
	frameTime := newTime - lastTime
	const maxUpdate = .25
	if frameTime > maxUpdate {
		frameTime = maxUpdate
	}
	lastTime = newTime
	accumulator += frameTime

	dt := 1. / 180
	for accumulator >= dt {
		space.Space.Step(dt)
		accumulator -= dt
	}
}
