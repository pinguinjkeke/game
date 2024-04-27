package renderer

import (
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"golang.org/x/image/colornames"
)

func RenderTerrain(ecs *ecs.ECS, screen *ebiten.Image) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	component.Terrain.Each(ecs.World, func(terrainEntry *donburi.Entry) {
		terrain := component.Terrain.Get(terrainEntry)
		shape := component.Shape.Get(terrainEntry)
		position := shape.Body().Position()

		// TODO add camera visibility check

		imageOptions := &ebiten.DrawImageOptions{}
		camera.GetTranslation(imageOptions, position.X-terrain.StartX, position.Y-terrain.StartY)
		screen.DrawImage(terrain.Sprite, imageOptions)

		o := &ebiten.DrawImageOptions{}
		camera.GetTranslation(o, 0, position.Y)
		image := ebiten.NewImage(4, 4)
		image.Fill(colornames.Red)
		screen.DrawImage(image, o)
		o.GeoM.Reset()
		camera.GetTranslation(o, position.X-2, position.Y+1)
		screen.DrawImage(image, o)
		o.GeoM.Translate(position.X, 0)
		screen.DrawImage(image, o)
	})
}