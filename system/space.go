package system

import (
	"game/component"
	"github.com/yohamta/donburi/ecs"
)

func UpdateSpace(ecs *ecs.ECS) {
	spaceEntry := component.Space.MustFirst(ecs.World)
	space := component.Space.Get(spaceEntry)

	for _, object := range space.Space.Objects() {
		object.Update()
	}
}
