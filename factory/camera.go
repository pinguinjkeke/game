package factory

import (
	"game/archetype"
	"game/component"
	camera "github.com/melonfunction/ebiten-camera"
	"github.com/yohamta/donburi/ecs"
)

func CreateCamera(ecs *ecs.ECS, levelWidth, levelHeight int) {
	cameraEntry := archetype.Camera.Spawn(ecs)

	ebitenCamera := camera.NewCamera(levelWidth, levelHeight, 0, 0, 0, 1)

	component.Camera.Set(cameraEntry, ebitenCamera)
}
