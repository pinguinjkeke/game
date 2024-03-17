package component

import (
	"github.com/melonfunction/ebiten-camera"
	"github.com/yohamta/donburi"
)

var Camera = donburi.NewComponentType[camera.Camera]()
