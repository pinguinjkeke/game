package renderer

import (
	"game/component"
	"game/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"image/color"
)

var wallColor = color.RGBA{R: 60, G: 60, B: 60, A: 255}

func RenderWall(ecs *ecs.ECS, screen *ebiten.Image) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	tags.Wall.Each(ecs.World, func(entry *donburi.Entry) {
		object := component.Object.Get(entry)

		vector.DrawFilledRect(
			screen,
			float32(object.Position.X),
			float32(object.Position.Y),
			float32(object.Size.X),
			float32(object.Size.Y),
			wallColor,
			false,
		)

		o := &ebiten.DrawImageOptions{}
		camera.GetTranslation(o, object.Position.X, object.Position.Y)
		image := ebiten.NewImage(int(object.Size.X), int(object.Size.Y))
		image.Fill(wallColor)
		screen.DrawImage(image, o)
	})
}
