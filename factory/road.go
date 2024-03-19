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
	marksCount := camera.Surface.Bounds().Dx() / markWidth
	marks := make([]float64, marksCount)

	for i, _ := range marks {
		marks[i] = float64(camera.Width*i) / 2
	}

	image := ebiten.NewImage(markWidth, markHeight)
	image.Fill(markColor)

	component.Road.Set(roadEntry, &component.RoadData{
		MarksX:     marks,
		MarkWidth:  160,
		MarkHeight: markHeight,
		MarksImage: image,
	})
}
