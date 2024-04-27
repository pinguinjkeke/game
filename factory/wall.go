package factory

import (
	"game/archetype"
	"game/component"
	"github.com/jakecoffman/cp/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateWall(ecs *ecs.ECS, x, y, width, height float64) *donburi.Entry {
	spaceEntry := component.Space.MustFirst(ecs.World)
	space := component.Space.Get(spaceEntry)

	position := cp.Vector{X: x, Y: y}
	shape := space.Space.AddShape(
		cp.NewSegment(
			space.Space.StaticBody,
			position,
			position.Add(cp.Vector{X: width, Y: height}),
			0,
		),
	)
	shape.SetElasticity(1)
	shape.SetFriction(1)

	wallEntry := archetype.Wall.Spawn(ecs)
	component.Shape.Set(wallEntry, &component.ShapeData{
		Shape:  shape,
		Width:  width,
		Height: height,
	})

	return wallEntry
}
