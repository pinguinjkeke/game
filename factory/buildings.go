package factory

import (
	"game/archetype"
	"game/assets"
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi/ecs"
	"image/color"
	"math"
	"math/rand"
)

const (
	buildingMarginX   = 8
	windowSize        = 2
	minWindows        = 6
	minWindowBlocks   = 2
	windowColumns     = 6
	windowOffset      = windowSize
	windowBlockOffset = 6
	buildingPaddingY  = 5
)

var buildingColors = []color.Color{
	color.RGBA{R: 0x23, G: 0x2E, B: 0x32, A: 0xff},
	color.RGBA{R: 0x30, G: 0x3A, B: 0x43, A: 0xff},
}

var lightsOnColors = []color.Color{
	color.RGBA{R: 0xB2, G: 0xB5, B: 0xB7, A: 0xff},
	color.RGBA{R: 0xB2, G: 0x9F, B: 0x8A, A: 0xff},
}
var lightsOffColors = []color.Color{
	color.RGBA{R: 0x30, G: 0x3A, B: 0x43, A: 0xff},
	color.RGBA{R: 0x41, G: 0x4A, B: 0x52, A: 0xff},
}

var lampSprite *ebiten.Image
var antennaSprite *ebiten.Image

func CreateBuildings(ecs *ecs.ECS, levelWidth float64) {
	initSprites()

	buildingsEntry := archetype.Buildings.Spawn(ecs)
	offset := 0.0
	buildings := make([]*component.BuildingData, 0, 10)

	for offset < levelWidth {
		building := createBuilding(offset)
		buildings = append(buildings, building)

		offset += float64(buildingMarginX*(rand.Intn(8)-3)) + float64(building.Sprite.Bounds().Dx())
	}

	component.Buildings.Set(buildingsEntry, &component.BuildingsData{Buildings: buildings})
}

func createBuilding(positionX float64) *component.BuildingData {
	doubledWindow := rand.Intn(4) == 1
	windowWidth, buildingMarginX := windowSize, windowOffset
	windowRows := minWindows + rand.Intn(5)
	windowBlocks := minWindowBlocks + rand.Intn(5)
	layer := rand.Intn(2)

	if doubledWindow {
		windowWidth, buildingMarginX = windowWidth*2, windowWidth
		windowRows = windowRows / 2
	}

	width := buildingMarginX*2 + windowRows*windowWidth + (windowRows-1)*windowOffset
	height := windowBlocks * (windowColumns*windowSize + (windowColumns-1)*windowOffset)
	height += windowBlocks * windowBlockOffset

	buildingSprite := ebiten.NewImage(width, height)
	buildingSprite.Fill(buildingColors[layer])

	x := buildingMarginX
	windows := make([]*component.BuildingWindowData, 0, 20)

	for x < width {
		y := buildingPaddingY

		for i := 0; i < windowBlocks; i++ {
			if y >= height {
				break
			}

			for j := 0; j < windowColumns; j++ {
				windows = append(windows, &component.BuildingWindowData{
					X:      x,
					Y:      y,
					Lights: rand.Intn(4) == 1,
				})

				y += windowSize + windowOffset
			}

			y += buildingPaddingY
		}

		x += windowWidth + windowOffset
	}

	sprite, oX, oY := addLights(buildingSprite)

	buildingData := &component.BuildingData{
		X:           positionX,
		OffsetX:     oX,
		OffsetY:     oY,
		Layer:       layer,
		Sprite:      sprite,
		Windows:     windows,
		WindowWidth: windowWidth,
	}

	for i := 0; i < len(windows); i++ {
		RenderLights(buildingData, i)
	}

	return buildingData
}

func RenderLights(building *component.BuildingData, windowIndex int) {
	window := building.Windows[windowIndex]
	color := lightsOnColors[building.Layer]

	if !window.Lights {
		color = lightsOffColors[building.Layer]
	}

	vector.DrawFilledRect(
		building.Sprite,
		float32(building.OffsetX+float64(window.X)),
		float32(building.OffsetY+float64(window.Y)),
		float32(building.WindowWidth),
		float32(windowSize),
		color,
		false,
	)

	if window.Lights {
		vector.DrawFilledRect(
			building.Sprite,
			float32(building.OffsetX+float64(window.X+building.WindowWidth-1)),
			float32(building.OffsetY+float64(window.Y+windowSize-1)),
			1,
			1,
			lightsOffColors[building.Layer],
			false,
		)
	}
}

func addLights(image *ebiten.Image) (newImage *ebiten.Image, oX, oY float64) {
	oX, oY = float64(lampSprite.Bounds().Dx()), 2.0
	drawCrane := rand.Intn(15) == 0
	drawAntenna := false

	if drawCrane {
		scaleFactor := float64(image.Bounds().Dx()) / float64(assets.CraneSprite.Bounds().Dx())
		oX, oY = 0, float64(assets.CraneSprite.Bounds().Dy())*scaleFactor
	} else if !drawCrane && rand.Intn(10) == 0 {
		drawAntenna = true
		oY = float64(antennaSprite.Bounds().Dy())
	}

	newImage = ebiten.NewImage(
		image.Bounds().Dx()+int(math.Ceil(oX)),
		image.Bounds().Dy()+int(math.Ceil(oY)),
	)

	if drawCrane {
		options := &ebiten.DrawImageOptions{}
		scaleFactor := float64(image.Bounds().Dx()) / float64(assets.CraneSprite.Bounds().Dx())
		options.GeoM.Scale(scaleFactor, scaleFactor)
		newImage.DrawImage(assets.CraneSprite, options)
	} else if drawAntenna {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(image.Bounds().Dx()/2-2), 0)
		newImage.DrawImage(antennaSprite, options)
	}

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(oX), float64(oY))
	newImage.DrawImage(image, options)

	if !drawCrane {
		options = &ebiten.DrawImageOptions{}
		options.GeoM.Translate(2, oY-2)
		newImage.DrawImage(lampSprite, options)
		options.GeoM.Translate(float64(image.Bounds().Dx())-2, 0)
		newImage.DrawImage(lampSprite, options)
	}

	return newImage, oX, oY
}

func initSprites() {
	if antennaSprite != nil && lampSprite != nil {
		return
	}

	antennaSprite, lampSprite = ebiten.NewImage(5, 30), ebiten.NewImage(5, 5)
	lampSprite.Fill(color.RGBA{R: 0xA3, G: 0x6C, B: 0x7B, A: 0xaa})
	antennaSprite.DrawImage(lampSprite, &ebiten.DrawImageOptions{})
	vector.DrawFilledRect(
		antennaSprite,
		float32(2),
		float32(5),
		float32(1),
		float32(25),
		buildingColors[1],
		false,
	)
}
