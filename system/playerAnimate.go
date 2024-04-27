package system

import (
	"game/component"
	"game/physics"
	"github.com/jakecoffman/cp/v2"
	"github.com/solarlune/ebitick"
	"github.com/yohamta/donburi/ecs"
	"math"
	"time"
)

func UpdatePlayerAnimation(ecs *ecs.ECS) {
	playerEntry := component.Player.MustFirst(ecs.World)
	player := component.Player.Get(playerEntry)

	timerEntry := component.Timer.MustFirst(ecs.World)
	timer := component.Timer.Get(timerEntry)

	playerShape := component.Shape.Get(playerEntry)
	velocity := playerShape.Shape.Body().Velocity()

	animations := component.Animation.Get(playerEntry)
	animations.GetActive().Update()

	velocityX := math.Abs(velocity.X)

	if !player.Grounded && velocity.Y != 0 {
		jump(animations, velocity)

		animations.PauseAtStart(component.PlayerRunningLandingAnimation)

		return
	}

	animations.PauseAtStart(
		component.PlayerJumpAnimation,
		component.PlayerFallAnimation,
		component.PlayerRunningJumpAnimation,
		component.PlayerRunningFallAnimation,
	)

	if velocityX != 0 {
		move(animations, player, velocity)
	} else {
		stand(animations, timer)
	}
}

func jump(animations *component.AnimationData, velocity cp.Vector) {
	animations.CancelStandingTimer()
	running := math.Abs(velocity.X) > physics.MaxWalkingSpeed

	if velocity.Y < 0 {
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

func move(animations *component.AnimationData, player *component.PlayerData, velocity cp.Vector) {
	animations.CancelStandingTimer()

	speed := math.Abs(velocity.X)

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
