package system

import (
	"game/component"
	"game/physics"
	"github.com/solarlune/ebitick"
	"github.com/yohamta/donburi/ecs"
	"math"
	"time"
)

func UpdatePlayerAnimation(ecs *ecs.ECS) {
	playerEntry := component.Player.MustFirst(ecs.World)
	player := component.Player.Get(playerEntry)
	animations := component.Animation.Get(playerEntry)

	timerEntry := component.Timer.MustFirst(ecs.World)
	timer := component.Timer.Get(timerEntry)

	animations.GetActive().Update()

	speed := math.Abs(player.SpeedX)

	if player.SpeedY != 0 {
		jump(animations, player)

		animations.PauseAtStart(component.PlayerRunningLandingAnimation)

		return
	}

	animations.PauseAtStart(
		component.PlayerJumpAnimation,
		component.PlayerFallAnimation,
		component.PlayerRunningJumpAnimation,
		component.PlayerRunningFallAnimation,
	)

	if speed != 0 {
		move(animations, player)
	} else {
		stand(animations, timer)
	}
}

func jump(animations *component.AnimationData, player *component.PlayerData) {
	animations.CancelStandingTimer()
	running := math.Abs(player.SpeedX) > physics.MaxWalkingSpeed

	if player.SpeedY < 0 {
		animations.Resume(component.PlayerJumpAnimation, component.PlayerRunningJumpAnimation)
		animations.Active = component.PlayerJumpAnimation

		if running {
			animations.Active = component.PlayerRunningJumpAnimation
		}
	} else {
		animations.Resume(component.PlayerFallAnimation, component.PlayerRunningFallAnimation)
		animations.Active = component.PlayerFallAnimation

		if running {
			animations.Active = component.PlayerRunningFallAnimation
		}
	}
}

func stand(animations *component.AnimationData, timer *ebitick.TimerSystem) {
	animations.Active = component.PlayerStandingAnimation

	if animations.StandingTimer == nil {
		animations.PauseAtStart(component.PlayerStandingAnimation)

		animations.StandingTimer = timer.After(4*time.Second, func() {
			animations.Animations[component.PlayerStandingAnimation].Resume()
		})
	}
}

func move(animations *component.AnimationData, player *component.PlayerData) {
	animations.CancelStandingTimer()

	speed := math.Abs(player.SpeedX)

	if player.JustChangedMovingDirection {
		animations.Reset(component.PlayerRunAnimation)
	}

	if speed > physics.MaxWalkingSpeed {
		animations.ActivateAndResume(component.PlayerRunAnimation)

		if !player.Running {
			animations.ActivateAndResume(component.PlayerStopRunAnimation)
		}

		if player.JustStoppedRunning {
			animations.PauseAtStart(component.PlayerRunAnimation)
			animations.PauseAtStart(component.PlayerStopRunAnimation)
		}

		if player.JustLanded {
			animations.ActivateAndResume(component.PlayerRunningLandingAnimation)
		}
	} else {
		animations.ActivateAndResume(component.PlayerWalkAnimation)
	}
}
