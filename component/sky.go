package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SkyData struct {
	Layers []*ebiten.Image
	MoonX  int
	MoonY  int
}

var Sky = donburi.NewComponentType[SkyData]()
