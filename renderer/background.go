package renderer

import (
	"game/assets"
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/melonfunction/ebiten-camera"
	"github.com/yohamta/donburi/ecs"
)

const BackgroundSize = 225

func RenderBackground(ecs *ecs.ECS, screen *ebiten.Image) {
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	layers := [2]*ebiten.Image{
		ebiten.NewImage(camera.Surface.Bounds().Dx(), camera.Surface.Bounds().Dy()),
		ebiten.NewImage(camera.Surface.Bounds().Dx(), camera.Surface.Bounds().Dy()),
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

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(-camera.X*0.1, 0)
	layers[0].DrawImage(buildings.Layers[0], options)
	options.GeoM.Translate(-camera.X*0.02, 0)
	layers[1].DrawImage(buildings.Layers[1], options)
}

func renderClouds(ecs *ecs.ECS, layers [2]*ebiten.Image, camera *camera.Camera) {
	skyEntry := component.Sky.MustFirst(ecs.World)
	sky := component.Sky.Get(skyEntry)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(camera.X*0.03, 0)
	layers[0].DrawImage(sky.Layers[0], options)
	options.GeoM.Translate(camera.X*0.01, 0)
	layers[1].DrawImage(sky.Layers[1], options)
}
