package factory

import (
	"game/archetype"
	"game/component"
	camera "github.com/melonfunction/ebiten-camera"
	"github.com/yohamta/donburi/ecs"
)

func CreateCamera(ecs *ecs.ECS, windowWidth, windowHeight int) {
	cameraEntry := archetype.Camera.Spawn(ecs)

	ebitenCamera := camera.NewCamera(windowWidth, windowHeight, 0, 0, 0, 1)

	component.Camera.Set(cameraEntry, ebitenCamera)
}
