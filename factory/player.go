package factory

import (
	"game/archetype"
	"game/component"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePlayer(ecs *ecs.ECS, x, y float64) *donburi.Entry {
	controlsEntry := component.Controls.MustFirst(ecs.World)
	controls := component.Controls.Get(controlsEntry)

	object := resolv.NewObject(x, y, component.PlayerFrameWidth/2, component.PlayerFrameHeight)
	object.SetShape(resolv.NewRectangle(0, 0, object.Size.X, object.Size.Y))

	playerEntry := archetype.Player.Spawn(ecs)

	component.Object.Set(playerEntry, object)
	component.ControlsHandler.Set(playerEntry, controls.NewHandler())
	component.Player.Set(playerEntry, &component.PlayerData{MovingDirection: 1})

	return playerEntry
}
