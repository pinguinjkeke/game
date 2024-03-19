package renderer

import (
	"game/assets"
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

func RenderRoad(ecs *ecs.ECS, screen *ebiten.Image) {
	roadEntry := component.Road.MustFirst(ecs.World)
	road := component.Road.Get(roadEntry)
	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(0, float64(screen.Bounds().Dy()-assets.RoadSprite.Bounds().Dy()))
	options.GeoM.Scale(float64(screen.Bounds().Dx()), 1)
	screen.DrawImage(assets.RoadSprite, options)

	markY := float64(screen.Bounds().Dy() - assets.RoadSprite.Bounds().Dy()/2 - road.MarkHeight/2)

	for _, markX := range road.MarksX {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Skew(0.8, 0)
		options.GeoM.Translate(markX-float64(int(camera.X)%640), markY)
		screen.DrawImage(road.MarksImage, options)
	}
}
