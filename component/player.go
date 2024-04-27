package component

import (
	"game/chipmunk"
	"github.com/yohamta/donburi"
)

const (
	PlayerFrameHeight     = 64
	PlayerFrameWidth      = 64
	PlayerWidth           = 22
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
	Grounded                   bool
	Jumping                    bool
	JustLanded                 bool
	landing                    playerLandingData
	LandDistance               float64
	MovingDirection            int
	JustChangedMovingDirection bool
	Running                    bool
	JustStoppedRunning         bool
	SpeedX                     float64
	SpeedY                     float64
}

type playerLandingData struct {
	startPosition   float64
	movingDirection int
}

func (p *PlayerData) Land(startPosition float64) {
	if p.Running {
		p.JustLanded = true
		p.landing.startPosition = startPosition
		p.landing.movingDirection = p.MovingDirection
	}
}

func (p *PlayerData) FinishLanding(position float64) {
	direction := float64(p.landing.movingDirection)

	if !p.Running ||
		p.landing.movingDirection != p.MovingDirection ||
		position*direction > p.landing.startPosition*direction+chipmunk.LandingDistance {
		p.landing.startPosition = 0
		p.landing.movingDirection = 0
		p.JustLanded = false

		return
	}
}

var Player = donburi.NewComponentType[PlayerData]()
