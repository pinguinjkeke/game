package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type RoadData struct {
	DrainsX    []float64
	DrainsY    float64
	MarksX     []float64
	MarkWidth  float64
	MarkHeight int
	MarksImage *ebiten.Image
	SewersX    []float64
	SewersY    float64
}

var Road = donburi.NewComponentType[RoadData]()
