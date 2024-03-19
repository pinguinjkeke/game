package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type RoadData struct {
	MarksX     []float64
	MarkWidth  float64
	MarkHeight int
	MarksImage *ebiten.Image
}

var Road = donburi.NewComponentType[RoadData]()
