package factory

import (
	"game/archetype"
	"game/component"
	"game/physics"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateWall(ecs *ecs.ECS, x, y, width, height float64) *donburi.Entry {
	object := resolv.NewObject(x, y, width, height, physics.TagSolid)
	wallEntry := archetype.Wall.Spawn(ecs)
	component.Object.Set(wallEntry, object)

	return wallEntry
}
