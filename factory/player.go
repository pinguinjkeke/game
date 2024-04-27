package factory

import (
	"game/archetype"
	"game/component"
	"game/physics"
	"github.com/jakecoffman/cp/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePlayer(ecs *ecs.ECS, x, y float64) *donburi.Entry {
	playerEntry := archetype.Player.Spawn(ecs)

	controlsEntry := component.Controls.MustFirst(ecs.World)
	controls := component.Controls.Get(controlsEntry)
	component.ControlsHandler.Set(playerEntry, controls.NewHandler())

	spaceEntry := component.Space.MustFirst(ecs.World)
	space := component.Space.Get(spaceEntry)

	body := space.Space.AddBody(cp.NewBody(1, cp.INFINITY))
	body.SetPosition(cp.Vector{X: x, Y: y})
	body.SetVelocityUpdateFunc(physics.PlayerUpdateVelocity(playerEntry))

	shape := space.Space.AddShape(cp.NewBox(body, component.PlayerWidth, component.PlayerFrameHeight, 0))
	shape.SetElasticity(0)
	shape.SetFriction(0)
	shape.SetCollisionType(1)

	component.Shape.Set(playerEntry, &component.ShapeData{
		Shape:  shape,
		Width:  component.PlayerWidth,
		Height: component.PlayerFrameHeight,
	})
	component.Player.Set(playerEntry, &component.PlayerData{MovingDirection: 1})

	return playerEntry
}
