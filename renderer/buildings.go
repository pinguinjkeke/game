package renderer

import (
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi/ecs"
	"image/color"
)

const borderWidth = 2

var borderColor = color.RGBA{R: 0x42, G: 0x4D, B: 0x5A, A: 0xff}

var buildingColors = []color.Color{
	color.RGBA{R: 0x27, G: 0x31, B: 0x3E, A: 0xff},
	color.RGBA{R: 0x32, G: 0x3E, B: 0x4F, A: 0xff},
}

var lightsOnColors = []color.Color{
	color.RGBA{R: 0xAE, G: 0xB3, B: 0x76, A: 0xff},
	color.RGBA{R: 0xF8, G: 0xFF, B: 0xA9, A: 0xff},
}
var lightsOffColors = []color.Color{
	color.RGBA{R: 0x18, G: 0x1F, B: 0x27, A: 0xff},
	color.RGBA{R: 0x1E, G: 0x25, B: 0x2E, A: 0xff},
}

func RenderBuildings(ecs *ecs.ECS, screen *ebiten.Image) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	layers := [2]*ebiten.Image{
		ebiten.NewImage(camera.Width, camera.Height),
		ebiten.NewImage(camera.Width, camera.Height),
	}
	imageOptions := &ebiten.DrawImageOptions{}

	cameraStartX, cameraEndX := int(camera.X), int(camera.X)+camera.Width

	buildingsEntry := component.Buildings.MustFirst(ecs.World)
	buildings := component.Buildings.Get(buildingsEntry)

	for _, building := range buildings.Buildings {
		parallax := 1.0
		if building.Layer == 0 {
			parallax = 0.95
		}
		parallaxOffset := int(camera.X * parallax)

		startX := building.X - parallaxOffset
		cameraMinX, cameraMaxX := cameraStartX-parallaxOffset, int(cameraEndX+parallaxOffset)

		if cameraMinX > startX+parallaxOffset || cameraMaxX < startX+building.Width {
			continue
		}

		vector.DrawFilledRect(
			layers[building.Layer],
			float32(startX),
			float32(building.Y)-borderWidth,
			float32(building.Width),
			borderWidth,
			borderColor,
			false,
		)
		vector.DrawFilledRect(
			layers[building.Layer],
			float32(startX+building.Width),
			float32(building.Y)-borderWidth,
			borderWidth,
			float32(building.Height),
			borderColor,
			false,
		)
		vector.DrawFilledRect(
			layers[building.Layer],
			float32(startX),
			float32(building.Y),
			float32(building.Width),
			float32(building.Height+borderWidth),
			buildingColors[building.Layer],
			false,
		)

		for _, window := range building.Windows {
			windowColor := lightsOnColors[building.Layer]

			if !window.Light {
				windowColor = lightsOffColors[building.Layer]
			}

			vector.DrawFilledRect(
				layers[building.Layer],
				float32(startX+window.X),
				float32(building.Y+window.Y),
				float32(building.WindowWidth),
				float32(building.WindowHeight),
				windowColor,
				false,
			)
		}
	}

	for _, layer := range layers {
		camera.Surface.DrawImage(layer, imageOptions)
	}
}
