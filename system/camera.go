package system

import (
	"game/component"
	"github.com/yohamta/donburi/ecs"
)

func UpdateCamera(ecs *ecs.ECS) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)
	playerEntry := component.Player.MustFirst(ecs.World)
	playerShape := component.Shape.Get(playerEntry)

	camera.SetPosition(playerShape.Body().Position().X, 370)
}
