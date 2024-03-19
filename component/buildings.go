package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ebitick"
	"github.com/yohamta/donburi"
)

type BuildingsData struct {
	Buildings []*BuildingData
	Timer     *ebitick.Timer
}

type BuildingData struct {
	X           float64
	OffsetX     float64
	OffsetY     float64
	Layer       int
	Sprite      *ebiten.Image
	Windows     []*BuildingWindowData
	WindowWidth int
}

type BuildingWindowData struct {
	X      int
	Y      int
	Lights bool
}

var Buildings = donburi.NewComponentType[BuildingsData]()
