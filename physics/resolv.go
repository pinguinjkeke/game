package physics

import (
	"game/component"
	"github.com/yohamta/donburi"
)

const (
	TagSolid = "solid"
)

const (
	Friction                 = 0.5
	Acceleration             = 1.0
	RunningAcceleration      = 1.5
	MaxWalkingSpeed          = 1.5
	MaxRunningSpeed          = 4.0
	XSpeedDiff               = MaxRunningSpeed - MaxWalkingSpeed
	LandingDistance          = 280.0
	JumpSpeed                = 10.0
	Gravity                  = 0.75
	TopPlatformSlideDistance = 8
)

func Add(space *donburi.Entry, objects ...*donburi.Entry) {
	for _, obj := range objects {
		object := component.Object.Get(obj)

		component.Space.Get(space).Space.Add(object)
	}
}
