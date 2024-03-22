package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SkyData struct {
	Layers     []*ebiten.Image
	LayerWidth float64
	X0         float64
	X1         float64
	MoonX      int
	MoonY      int
}

var Sky = donburi.NewComponentType[SkyData]()
