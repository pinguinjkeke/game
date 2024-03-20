package factory

import (
	"game/archetype"
	"game/assets"
	"game/component"
	"game/renderer"
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

type buildingData struct {
	X           float64
	Layer       int
	Sprite      *ebiten.Image
	WindowWidth int
}

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

func CreateBuildings(ecs *ecs.ECS) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)
	backgroundWidth := camera.Surface.Bounds().Dx() * 2

	initSprites()

	layers := []*ebiten.Image{
		ebiten.NewImage(backgroundWidth, renderer.BackgroundSize),
		ebiten.NewImage(backgroundWidth, renderer.BackgroundSize),
	}

	buildingsEntry := archetype.Buildings.Spawn(ecs)
	offset := 0.0
	windows := make([]*component.BuildingWindowData, 0, 100)

	for offset < float64(backgroundWidth) {
		building, buildingWindows := createBuilding(offset)
		windows = append(windows, buildingWindows...)

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(building.X, float64(renderer.BackgroundSize-building.Sprite.Bounds().Dy()))
		layers[building.Layer].DrawImage(building.Sprite, options)

		offset += float64(buildingMarginX*(rand.Intn(8)-3)) + float64(building.Sprite.Bounds().Dx())
	}

	for _, window := range windows {
		RenderWindow(window, layers)
	}

	component.Buildings.Set(buildingsEntry, &component.BuildingsData{
		Layers:  layers,
		Windows: windows,
	})
}

func createBuilding(positionX float64) (building *buildingData, windows []*component.BuildingWindowData) {
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

	sprite, oX := addLights(buildingSprite)

	windows = make([]*component.BuildingWindowData, 0, 20)
	x := buildingMarginX
	for x < width {
		y := buildingPaddingY

		for i := 0; i < windowBlocks; i++ {
			if y >= height {
				break
			}

			for j := 0; j < windowColumns; j++ {
				windows = append(windows, &component.BuildingWindowData{
					X:      int(positionX+oX) + x,
					Y:      renderer.BackgroundSize - height + y,
					Layer:  layer,
					Lights: rand.Intn(4) == 1,
					Width:  windowWidth,
				})

				y += windowSize + windowOffset
			}

			y += buildingPaddingY
		}

		x += windowWidth + windowOffset
	}

	buildingData := &buildingData{
		X:           positionX,
		Layer:       layer,
		Sprite:      sprite,
		WindowWidth: windowWidth,
	}

	return buildingData, windows
}

func RenderWindow(window *component.BuildingWindowData, layers []*ebiten.Image) {
	color := lightsOnColors[window.Layer]

	if !window.Lights {
		color = lightsOffColors[window.Layer]
	}

	vector.DrawFilledRect(
		layers[window.Layer],
		float32(window.X),
		float32(window.Y),
		float32(window.Width),
		float32(windowSize),
		color,
		false,
	)

	if window.Lights {
		vector.DrawFilledRect(
			layers[window.Layer],
			float32(window.X+window.Width-1),
			float32(window.Y-1),
			1,
			1,
			lightsOffColors[window.Layer],
			false,
		)
	}
}

func addLights(image *ebiten.Image) (newImage *ebiten.Image, oX float64) {
	oX, oY := float64(lampSprite.Bounds().Dx()), 2.0
	drawCrane := rand.Intn(10) == 0
	drawAntenna := false

	if drawCrane {
		scaleFactor := float64(image.Bounds().Dx()) / float64(assets.CraneSprite.Bounds().Dx())
		oX, oY = 0, float64(assets.CraneSprite.Bounds().Dy())*scaleFactor
	} else if !drawCrane && rand.Intn(5) == 0 {
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
		options.GeoM.Translate(float64(image.Bounds().Dx()/2), 0)
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

	return newImage, oX
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
