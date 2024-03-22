package renderer

import (
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func RenderTerrain(ecs *ecs.ECS, screen *ebiten.Image) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	component.Terrain.Each(ecs.World, func(terrainEntry *donburi.Entry) {
		terrain := component.Terrain.Get(terrainEntry)
		object := component.Object.Get(terrainEntry)

		// TODO add camera visibility check

		imageOptions := &ebiten.DrawImageOptions{}
		camera.GetTranslation(imageOptions, object.Position.X-component.PlayerFrameWidth/4, object.Position.Y)
		screen.DrawImage(terrain.Sprite, imageOptions)
	})
}
