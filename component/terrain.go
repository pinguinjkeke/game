package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type TerrainData struct {
	Sprite *ebiten.Image
}

var Terrain = donburi.NewComponentType[TerrainData]()
