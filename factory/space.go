package factory

import (
	"game/archetype"
	"game/chipmunk"
	"game/component"
	"github.com/jakecoffman/cp/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateSpace(ecs *ecs.ECS, width, height int) *donburi.Entry {
	spaceEntry := archetype.Space.Spawn(ecs)

	space := cp.NewSpace()
	space.Iterations = 10
	space.SetGravity(cp.Vector{X: 0, Y: chipmunk.Gravity})

	component.Space.Set(spaceEntry, &component.SpaceData{
		WindowWidth:  width,
		WindowHeight: height,
		Space:        space,
	})

	return spaceEntry
}
