package factory

import (
	"game/assets"
	"game/component"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
	"time"
)

func CreatePlayerAnimation(ecs *ecs.ECS) {
	playerEntry := component.Player.MustFirst(ecs.World)

	timerEntry := component.Timer.MustFirst(ecs.World)
	timer := component.Timer.Get(timerEntry)

	g32 := ganim8.NewGrid(component.PlayerFrameWidth, component.PlayerFrameHeight, component.PlayerFrameWidth*16, component.PlayerFrameHeight*13)
	animations := []*ganim8.Animation{
		component.PlayerStandingAnimation: ganim8.New(
			assets.PlayerSprite,
			g32.Frames("1-16", 8, "1-16", 9, "1-16", 10, "1-16", 11, "1-16", 12, "1-6", 13),
			150*time.Millisecond,
		),
		component.PlayerWalkAnimation: ganim8.New(
			assets.PlayerSprite,
			g32.Frames("1-8", 1),
			component.PlayerAnimationTickMs*time.Millisecond,
		),
		component.PlayerRunAnimation: ganim8.New(
			assets.PlayerSprite,
			g32.Frames("1-9", 2),
			component.PlayerAnimationTickMs*time.Millisecond,
			func(anim *ganim8.Animation, loops int) {
				if loops > 0 {
					anim.GoToFrame(3)
				}
			},
		),
		component.PlayerStopRunAnimation: ganim8.New(
			assets.PlayerSprite,
			g32.Frames("10-15", 2),
			80*time.Millisecond,
		),
		component.PlayerJumpAnimation: ganim8.New(
			assets.PlayerSprite,
			g32.Frames("1-3", 3),
			80*time.Millisecond,
		),
		component.PlayerFallAnimation: ganim8.New(
			assets.PlayerSprite,
			g32.Frames("4-6", 3),
			100*time.Millisecond,
		),
		component.PlayerRunningJumpAnimation: ganim8.New(
			assets.PlayerSprite,
			g32.Frames("1-3", 4),
			80*time.Millisecond,
		),
		component.PlayerRunningFallAnimation: ganim8.New(
			assets.PlayerSprite,
			g32.Frames("4", 4),
			100*time.Millisecond,
		),
		component.PlayerRunningLandingAnimation: ganim8.New(
			assets.PlayerSprite,
			g32.Frames("5-13", 4),
			100*time.Millisecond,
		),
	}

	for _, animation := range animations {
		animation.PauseAtStart()
	}

	animationsData := &component.AnimationData{
		Active:     component.PlayerStandingAnimation,
		Animations: animations,
		StandingTimer: timer.After(4*time.Second, func() {
			animations[component.PlayerStandingAnimation].Resume()
		}),
	}

	component.Animation.Set(playerEntry, animationsData)
}
