package component

import (
	"github.com/jakecoffman/cp/v2"
	"github.com/yohamta/donburi"
)

type ShapeData struct {
	*cp.Shape
	Width  float64
	Height float64
}

var Shape = donburi.NewComponentType[ShapeData]()
