package renderer

import (
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func RenderPlayer(ecs *ecs.ECS, screen *ebiten.Image) {
	playerEntry := component.Player.MustFirst(ecs.World)
	player := component.Player.Get(playerEntry)
	animations := component.Animation.Get(playerEntry)
	object := component.Object.Get(playerEntry)
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	playerImage := ebiten.NewImage(component.PlayerFrameWidth, component.PlayerFrameHeight)

	animationOptions := &ganim8.DrawOptions{
		X:       component.PlayerFrameWidth / 2,
		Y:       component.PlayerFrameHeight / 2,
		ScaleX:  1,
		ScaleY:  1,
		OriginX: 0.5,
		OriginY: 0.5,
	}

	if player.MovingDirection < 0 {
		animationOptions.ScaleX = -1
	}

	animations.GetActive().Draw(playerImage, animationOptions)

	imageOptions := &ebiten.DrawImageOptions{}
	camera.GetTranslation(imageOptions, object.Position.X-object.Size.X, object.Position.Y)
	screen.DrawImage(playerImage, imageOptions)
}
