package factory

import (
	"game/archetype"
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
	"image/color"
)

const (
	markWidth  = 160
	markHeight = 6
)

var markColor = color.RGBA{R: 0xDE, G: 0xD7, B: 0xA7, A: 0xff}

func CreateRoad(ecs *ecs.ECS) {
	roadEntry := archetype.Road.Spawn(ecs)

	cameraEntry := component.Camera.MustFirst(ecs.World)
	camera := component.Camera.Get(cameraEntry)

	drains := make([]float64, 4)

	for i, _ := range drains {
		drains[i] = float64(i * camera.Width / 2)
	}

	marksCount := camera.Surface.Bounds().Dx() / markWidth
	marks := make([]float64, marksCount)

	for i, _ := range marks {
		marks[i] = float64(camera.Width*i) / 2
	}

	marksImage := ebiten.NewImage(markWidth, markHeight)
	marksImage.Fill(markColor)

	sewers := make([]float64, 4)
	for i, _ := range sewers {
		sewers[i] = float64(i*camera.Width/2 + camera.Width/4)
	}

	component.Road.Set(roadEntry, &component.RoadData{
		DrainsX:    drains,
		DrainsY:    118,
		MarksX:     marks,
		MarkWidth:  160,
		MarkHeight: markHeight,
		MarksImage: marksImage,
		SewersX:    sewers,
		SewersY:    90,
	})
}
