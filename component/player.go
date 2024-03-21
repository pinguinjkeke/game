package component

import (
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

const (
	PlayerFrameHeight     = 64
	PlayerFrameWidth      = 64
	PlayerAnimationTickMs = 140
)

const (
	PlayerStandingAnimation = iota
	PlayerWalkAnimation
	PlayerRunAnimation
	PlayerStopRunAnimation
	PlayerJumpAnimation
	PlayerFallAnimation
	PlayerRunningJumpAnimation
	PlayerRunningFallAnimation
	PlayerRunningLandingAnimation
)

type PlayerData struct {
	Ground                     *resolv.Object
	Jumping                    bool
	JustLanded                 bool
	LandDistance               float64
	MovingDirection            int
	JustChangedMovingDirection bool
	Running                    bool
	JustStoppedRunning         bool
	SpeedX                     float64
	SpeedY                     float64
}

var Player = donburi.NewComponentType[PlayerData]()
