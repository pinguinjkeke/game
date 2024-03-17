package renderer

import (
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

func RenderCameraClear(ecs *ecs.ECS, screen *ebiten.Image) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	camera.Surface.Clear()
}

func RenderCamera(ecs *ecs.ECS, screen *ebiten.Image) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	camera.Blit(screen)
}
