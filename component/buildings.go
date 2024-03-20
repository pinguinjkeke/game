package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ebitick"
	"github.com/yohamta/donburi"
)

type BuildingsData struct {
	Timer   *ebitick.Timer
	Layers  []*ebiten.Image
	Windows []*BuildingWindowData
}

type BuildingWindowData struct {
	X      int
	Y      int
	Width  int
	Layer  int
	Lights bool
}

var Buildings = donburi.NewComponentType[BuildingsData]()
