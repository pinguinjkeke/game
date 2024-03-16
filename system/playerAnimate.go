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

		return
	}

	animations.PauseAtStart(component.PlayerJumpAnimation)
	animations.PauseAtStart(component.PlayerFallAnimation)

	if speed != 0 {
		move(animations, player)
	} else {
		stand(animations, timer)
	}
}

func jump(animations *component.AnimationData, player *component.PlayerData) {
	animations.CancelStandingTimer()

	if player.SpeedY < 0 {
		animations.ActivateAndResume(component.PlayerJumpAnimation)
	} else {
		animations.ActivateAndResume(component.PlayerFallAnimation)
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
	speedDelta := physics.MaxRunningSpeed - physics.MaxWalkingSpeed

	if speed > physics.MaxWalkingSpeed {
		animation := animations.ActivateAndResume(component.PlayerRunAnimation)

		if !player.Running {
			animation = animations.ActivateAndResume(component.PlayerStopRunAnimation)
		}

		if speed < physics.MaxRunningSpeed-speedDelta/4 {
			animation.GoToFrame(2)
		} else if speed < physics.MaxRunningSpeed-speedDelta/2 {
			animation.GoToFrame(1)
		}
	} else {
		animations.ActivateAndResume(component.PlayerWalkAnimation)
	}
}
