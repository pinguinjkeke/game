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

	g32 := ganim8.NewGrid(component.PlayerFrameWidth, component.PlayerFrameHeight, component.PlayerFrameWidth*50, component.PlayerFrameHeight*7)
	animationDuration := component.PlayerAnimationTickMs * time.Millisecond
	animations := []*ganim8.Animation{
		component.PlayerStandingAnimation: ganim8.New(assets.PlayerSprite, g32.Frames("1-15", 1, "1-15", 2, "1-15", 3, "1-5", 4), 250*time.Millisecond),
		component.PlayerWalkAnimation:     ganim8.New(assets.PlayerSprite, g32.Frames("1-8", 5), animationDuration),
		component.PlayerRunAnimation: ganim8.New(assets.PlayerSprite, g32.Frames("1-9", 6), animationDuration, func(anim *ganim8.Animation, loops int) {
			if loops > 0 {
				anim.GoToFrame(3)
			}
		}),
		component.PlayerStopRunAnimation: ganim8.New(assets.PlayerSprite, g32.Frames("10-11", 6), 100*time.Millisecond),
		component.PlayerJumpAnimation:    ganim8.New(assets.PlayerSprite, g32.Frames("1-3", 7), 80*time.Millisecond),
		component.PlayerFallAnimation:    ganim8.New(assets.PlayerSprite, g32.Frames("3-6", 7), 100*time.Millisecond),
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
