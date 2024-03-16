package factory

import (
	"game/archetype"
	"game/component"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateSpace(ecs *ecs.ECS, width, height int) *donburi.Entry {
	spaceEntry := archetype.Space.Spawn(ecs)

	space := resolv.NewSpace(width, height, 16, 16)

	component.Space.Set(spaceEntry, &component.SpaceData{
		WindowWidth:  width,
		WindowHeight: height,
		Space:        space,
	})

	return spaceEntry
}
