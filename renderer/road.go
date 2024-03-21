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

	dY := float64(screen.Bounds().Dy() - assets.RoadSprite.Bounds().Dy())
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(0, dY)
	options.GeoM.Scale(float64(screen.Bounds().Dx()), 1)
	screen.DrawImage(assets.RoadSprite, options)

	markY := float64(screen.Bounds().Dy() - assets.RoadSprite.Bounds().Dy()/4 + road.MarkHeight)
	for _, markX := range road.MarksX {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Skew(0.8, 0)
		options.GeoM.Translate(markX-float64(int(camera.X)%640), markY)
		screen.DrawImage(road.MarksImage, options)
	}

	screenWidth := float64(camera.Surface.Bounds().Dx())
	cameraOffsetX := screenWidth - float64(int(camera.X)%int(screenWidth))
	for _, sewerX := range road.SewersX {
		options = &ebiten.DrawImageOptions{}

		offset := sewerX + cameraOffsetX

		if offset+float64(assets.SewerSprite.Bounds().Dx()) > screenWidth {
			offset -= screenWidth
		}

		options.GeoM.Translate(offset, dY+road.SewersY)
		screen.DrawImage(assets.SewerSprite, options)
	}

	for _, drainX := range road.DrainsX {
		options = &ebiten.DrawImageOptions{}

		offset := drainX + cameraOffsetX

		if offset+float64(assets.DrainSprite.Bounds().Dx()) > screenWidth {
			offset -= screenWidth
		}

		options.GeoM.Translate(offset, dY+road.DrainsY)
		screen.DrawImage(assets.DrainSprite, options)
	}
}
