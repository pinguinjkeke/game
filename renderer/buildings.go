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

var parallaxes = [...]float64{0.95, 1.0}

func RenderBuildings(ecs *ecs.ECS, screen *ebiten.Image) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	layers := [...]*ebiten.Image{
		ebiten.NewImage(camera.Width, camera.Height),
		ebiten.NewImage(camera.Width, camera.Height),
	}
	imageOptions := &ebiten.DrawImageOptions{}

	buildingsEntry := component.Buildings.MustFirst(ecs.World)
	buildings := component.Buildings.Get(buildingsEntry)

	for _, building := range buildings.Buildings {
		parallax := parallaxes[building.Layer]
		startX := int(float64(building.X) - camera.X*parallax)

		if startX+building.Width < 0 || startX > camera.Width {
			continue
		}

		vector.DrawFilledRect(
			layers[building.Layer],
			float32(startX),
			float32(building.Y),
			float32(building.Width),
			float32(building.Height+borderWidth),
			buildingColors[building.Layer],
			false,
		)

		vector.DrawFilledRect(
			layers[building.Layer],
			float32(startX),
			float32(building.Y)-borderWidth,
			float32(building.Width),
			borderWidth,
			borderColor,
			false,
		)

		if startX+building.Width <= camera.Width {
			vector.DrawFilledRect(
				layers[building.Layer],
				float32(startX+building.Width),
				float32(building.Y)-borderWidth,
				borderWidth,
				float32(building.Height),
				borderColor,
				false,
			)
		}

		vector.DrawFilledRect(
			layers[building.Layer],
			float32(startX+building.WindowOffsetX),
			float32(building.Y+building.WindowOffsetY),
			float32(building.Width-building.WindowOffsetX*2),
			float32(building.Height-building.WindowOffsetY*2),
			lightsOnColors[building.Layer],
			false,
		)

		for i := building.WindowWidth + building.WindowOffsetX; i < building.Width; i += building.WindowWidth {
			if startX+i > camera.Width {
				continue
			}

			vector.DrawFilledRect(
				layers[building.Layer],
				float32(startX+i),
				float32(building.Y+building.WindowOffsetY),
				float32(building.WindowOffsetX),
				float32(building.Height-building.WindowOffsetY*2),
				buildingColors[building.Layer],
				false,
			)

			i += building.WindowOffsetX
		}

		for i := building.WindowHeight + building.WindowOffsetY; i < building.Height; i += building.WindowHeight {
			vector.DrawFilledRect(
				layers[building.Layer],
				float32(startX+building.WindowOffsetX),
				float32(building.Y+i),
				float32(building.Width-building.WindowOffsetX*2),
				float32(building.WindowOffsetY),
				buildingColors[building.Layer],
				false,
			)

			i += building.WindowOffsetY
		}

		window := 0

		for i := 0; i < building.WindowRows; i++ {
			for j := 0; j < building.WindowColumns; j++ {
				windowX := (building.WindowOffsetX * (i + 1)) + building.WindowWidth*i

				if building.WindowLights[window] || startX+windowX > camera.Width {
					window++
					continue
				}

				windowY := (building.WindowOffsetY * (j + 1)) + building.WindowHeight*j

				vector.DrawFilledRect(
					layers[building.Layer],
					float32(startX+windowX),
					float32(building.Y+windowY),
					float32(building.WindowWidth),
					float32(building.WindowHeight),
					lightsOffColors[building.Layer],
					false,
				)

				window++
			}
		}
	}

	for _, layer := range layers {
		camera.Surface.DrawImage(layer, imageOptions)
	}
}
