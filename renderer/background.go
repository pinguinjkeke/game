package renderer

import (
	"game/assets"
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	camera "github.com/melonfunction/ebiten-camera"
	"github.com/yohamta/donburi/ecs"
)

const BackgroundSize = 225

var buildingParallaxes = [...]float64{0.04, 0.06}
var cloudParallaxes = [...]float64{0.002, 0.004}

func RenderBackground(ecs *ecs.ECS, screen *ebiten.Image) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	layers := [...]*ebiten.Image{
		ebiten.NewImage(camera.Width, BackgroundSize),
		ebiten.NewImage(camera.Width, BackgroundSize),
	}

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(float64(camera.Surface.Bounds().Dx()), 1)
	layers[0].DrawImage(assets.SkySprite, options)

	renderMoon(ecs, layers[0], camera)
	renderBuildings(ecs, layers, camera)
	renderClouds(ecs, layers, camera)

	for _, layer := range layers {
		camera.Surface.DrawImage(layer, &ebiten.DrawImageOptions{})
	}
}

func renderMoon(ecs *ecs.ECS, image *ebiten.Image, camera *camera.Camera) {
	skyEntry := component.Sky.MustFirst(ecs.World)
	sky := component.Sky.Get(skyEntry)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(0.5, 0.5)
	options.GeoM.Translate(float64(sky.MoonX), float64(sky.MoonY))

	image.DrawImage(assets.MoonSprite, options)
}

func renderBuildings(ecs *ecs.ECS, layers [2]*ebiten.Image, camera *camera.Camera) {
	buildingsEntry := component.Buildings.MustFirst(ecs.World)
	buildings := component.Buildings.Get(buildingsEntry)

	for _, building := range buildings.Buildings {
		parallax := buildingParallaxes[building.Layer]
		startX := building.X - camera.X*parallax

		if startX+float64(building.Sprite.Bounds().Dx()) < 0 || startX > float64(camera.Width) {
			continue
		}

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(startX, float64(layers[building.Layer].Bounds().Dy()-building.Sprite.Bounds().Dy()))
		layers[building.Layer].DrawImage(building.Sprite, options)
	}
}

func renderClouds(ecs *ecs.ECS, layers [2]*ebiten.Image, camera *camera.Camera) {
	skyEntry := component.Sky.MustFirst(ecs.World)
	sky := component.Sky.Get(skyEntry)

	for _, cloud := range sky.Clouds {
		parallax := cloudParallaxes[cloud.Layer]
		startX := int(camera.X*parallax - float64(cloud.X))

		if startX+cloud.Sprite.Bounds().Dx()*2 < 0 || startX > camera.Width {
			continue
		}

		options := &ebiten.DrawImageOptions{}

		if cloud.Flip {
			options.GeoM.Scale(-1, 1)
			options.GeoM.Translate(float64(cloud.Sprite.Bounds().Dx()), 0)
		}

		options.GeoM.Translate(float64(cloud.Sprite.Bounds().Dx()), 0)

		options.GeoM.Translate(float64(startX), float64(cloud.Y))
		options.GeoM.Scale(0.8, 0.8)

		layers[cloud.Layer].DrawImage(cloud.Sprite, options)
	}
}
