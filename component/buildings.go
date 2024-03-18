package component

import (
	"github.com/solarlune/ebitick"
	"github.com/yohamta/donburi"
)

type BuildingWindowData struct {
	X     int
	Y     int
	Light bool
}

type BuildingData struct {
	X             int
	Y             int
	Width         int
	Height        int
	WindowWidth   int
	WindowHeight  int
	WindowRows    int
	WindowColumns int
	WindowOffsetX int
	WindowOffsetY int
	WindowLights  []bool
	Layer         int
}

type BuildingsData struct {
	Buildings []*BuildingData
	Timer     *ebitick.Timer
}

var Buildings = donburi.NewComponentType[BuildingsData]()
