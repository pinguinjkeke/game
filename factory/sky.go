package factory

import (
	"game/archetype"
	"game/assets"
	"game/component"
	"game/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
	"image"
	"math"
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
	layerWidth := float64(backgroundWidth)

	layers := []*ebiten.Image{
		ebiten.NewImage(backgroundWidth, renderer.BackgroundSize),
		ebiten.NewImage(backgroundWidth, renderer.BackgroundSize),
	}

	cloudOffset, offset, options := windowWidth/8, 0, &ebiten.DrawImageOptions{}

	for offset+cloudStep*2 <= backgroundWidth {
		options.GeoM.Reset()

		if rand.Intn(3) == 1 {
			options.GeoM.Scale(-1, 1)
		}

		layer := rand.Intn(2)
		width := float64(sprites[0].Bounds().Dx())

		endCoordinate := math.Min(float64(offset)+width, layerWidth)

		options.GeoM.Translate(endCoordinate-width, float64(rand.Intn(renderer.BackgroundSize/3)))

		if rand.Intn(2) == 0 {
			layers[layer].DrawImage(sprites[0], options)
		}

		offset += cloudOffset - rand.Intn(cloudOffset/4)
	}

	offset = 0

	for offset+cloudStep*2 <= backgroundWidth {
		options.GeoM.Reset()

		if rand.Intn(3) == 1 {
			options.GeoM.Scale(-1, 1)
		}

		layer := rand.Intn(2)

		layers[layer].DrawImage(sprites[0], options)

		sprite := sprites[rand.Intn(len(sprites)-1)+1]
		width := float64(sprite.Bounds().Dx())

		endCoordinate := math.Min(float64(offset)+width, layerWidth)

		options.GeoM.Translate(endCoordinate-width, float64(rand.Intn(renderer.BackgroundSize/3)))

		layers[layer].DrawImage(sprite, options)
		offset += cloudOffset - rand.Intn(cloudOffset/4)
	}

	component.Sky.Set(skyEntry, &component.SkyData{
		Layers:     layers,
		LayerWidth: layerWidth,
		MoonX:      windowWidth - 40,
		MoonY:      20,
		X0:         -layerWidth,
		X1:         -layerWidth,
	})
}
