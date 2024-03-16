package component

import (
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

type SpaceData struct {
	Space        *resolv.Space
	WindowWidth  int
	WindowHeight int
}

func (s *SpaceData) Add(entry *donburi.Entry) {
	object := Object.Get(entry)

	s.Space.Add(object)
}

var Space = donburi.NewComponentType[SpaceData]()
