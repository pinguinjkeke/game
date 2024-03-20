package factory

import (
	"game/archetype"
	"game/assets"
	"game/component"
	"game/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
	"image"
	"math/rand"
)

const cloudStep = 64

var sprites = [...]*ebiten.Image{
	assets.CloudsSprite.SubImage(image.Rect(0, 0, cloudStep*2, cloudStep)).(*ebiten.Image),
	assets.CloudsSprite.SubImage(image.Rect(cloudStep*2, 0, cloudStep*2, cloudStep)).(*ebiten.Image),
	assets.CloudsSprite.SubImage(image.Rect(cloudStep*4, 0, cloudStep*2, cloudStep)).(*ebiten.Image),
	assets.CloudsSprite.SubImage(image.Rect(cloudStep*6, 0, cloudStep*2, cloudStep)).(*ebiten.Image),
}

func CreateSky(ecs *ecs.ECS) {
	skyEntry := archetype.Sky.Spawn(ecs)

	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)
	windowWidth := camera.Surface.Bounds().Dx()
	backgroundWidth := windowWidth * 2

	layers := []*ebiten.Image{
		ebiten.NewImage(backgroundWidth, renderer.BackgroundSize),
		ebiten.NewImage(backgroundWidth, renderer.BackgroundSize),
	}

	cloudOffset, offset, options := windowWidth/4, 0, &ebiten.DrawImageOptions{}

	for offset <= backgroundWidth {
		options.GeoM.Reset()
		options.GeoM.Translate(float64(offset), float64(rand.Intn(renderer.BackgroundSize/3)))

		if rand.Intn(3) == 1 {
			//options.GeoM.Scale(-1, 1)
		}

		layers[rand.Intn(2)].DrawImage(sprites[rand.Intn(len(sprites))], options)
		offset += cloudOffset - (rand.Intn(cloudOffset/4) + cloudOffset/4)
	}

	component.Sky.Set(skyEntry, &component.SkyData{
		Layers: layers,
		MoonX:  windowWidth - 40,
		MoonY:  20,
	})
}
