package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SkyData struct {
	Clouds []CloudData
	MoonX  int
	MoonY  int
}

type CloudData struct {
	X      int
	Y      int
	Layer  int
	Flip   bool
	Sprite *ebiten.Image
}

var Sky = donburi.NewComponentType[SkyData]()
