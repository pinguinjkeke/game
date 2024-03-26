package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type TerrainData struct {
	StartX float64
	StartY float64
	Sprite *ebiten.Image
}

var Terrain = donburi.NewComponentType[TerrainData]()
