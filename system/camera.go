package system

import (
	"game/component"
	"github.com/yohamta/donburi/ecs"
)

func UpdateCamera(ecs *ecs.ECS) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)
	playerEntry := component.Player.MustFirst(ecs.World)
	playerObject := component.Object.Get(playerEntry)

	camera.SetPosition(playerObject.Position.X, 320)
}
