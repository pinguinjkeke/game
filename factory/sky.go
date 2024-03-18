package factory

import (
	"game/archetype"
	"game/assets"
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
	"image"
	"math/rand"
)

const cloudStep = 64

var sprites = [...]*ebiten.Image{
	assets.CloudsSprite.SubImage(image.Rect(0, 0, cloudStep*7, cloudStep*2)).(*ebiten.Image),
	assets.CloudsSprite.SubImage(image.Rect(cloudStep*7, 0, cloudStep*11, cloudStep*2)).(*ebiten.Image),
	assets.CloudsSprite.SubImage(image.Rect(cloudStep*11, 0, cloudStep*17, cloudStep*2)).(*ebiten.Image),
	assets.CloudsSprite.SubImage(image.Rect(cloudStep*17, 0, cloudStep*21, cloudStep*2)).(*ebiten.Image),
	assets.CloudsSprite.SubImage(image.Rect(cloudStep*21, 0, cloudStep*23, cloudStep*2)).(*ebiten.Image),
}

func CreateSky(ecs *ecs.ECS, levelWidth, windowWidth int) {
	skyEntry := archetype.Sky.Spawn(ecs)

	cloudsPerWindow := float64(levelWidth/windowWidth) * 0.15
	clouds := make([]component.CloudData, int(cloudsPerWindow*2.0/(rand.Float64()/2)))
	windowOffset := windowWidth / len(clouds)

	for i, _ := range clouds {
		clouds[i] = component.CloudData{
			X:      i*windowOffset + rand.Intn(windowOffset/3),
			Y:      rand.Intn(120),
			Layer:  rand.Intn(2),
			Flip:   rand.Intn(2) == 1,
			Sprite: sprites[rand.Intn(len(sprites))],
		}
	}

	component.Sky.Set(skyEntry, &component.SkyData{
		Clouds: clouds,
		MoonX:  windowWidth - 60,
		MoonY:  40,
	})
}
