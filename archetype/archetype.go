package archetype

import (
	"game/layers"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type archetype struct {
	components []donburi.IComponentType
}

func newArchetype(components ...donburi.IComponentType) *archetype {
	return &archetype{
		components: components,
	}
}

func (a *archetype) Spawn(ecs *ecs.ECS, components ...donburi.IComponentType) *donburi.Entry {
	return ecs.World.Entry(
		ecs.Create(
			layers.Default,
			append(a.components, components...)...,
		),
	)
}
