package component

import (
	"github.com/jakecoffman/cp/v2"
	"github.com/yohamta/donburi"
)

type SpaceData struct {
	Space        *cp.Space
	WindowWidth  int
	WindowHeight int
}

var Space = donburi.NewComponentType[SpaceData]()
