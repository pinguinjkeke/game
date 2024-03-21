package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type BuildingsData struct {
	Layers []*ebiten.Image
}

var Buildings = donburi.NewComponentType[BuildingsData]()
