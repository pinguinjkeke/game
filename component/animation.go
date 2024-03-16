package component

import (
	"github.com/solarlune/ebitick"
	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

type AnimationData struct {
	Active        int
	Animations    []*ganim8.Animation
	StandingTimer *ebitick.Timer
}

func (a *AnimationData) Activate(animation int) *ganim8.Animation {
	a.Active = animation

	return a.Animations[animation]
}

func (a *AnimationData) ActivateAndResume(animation int) *ganim8.Animation {
	a.Animations[animation].Resume()
	a.Active = animation

	return a.Animations[animation]
}

func (a *AnimationData) GetActive() *ganim8.Animation {
	return a.Animations[a.Active]
}

func (a *AnimationData) PauseAtStart(animation int) {
	a.Animations[animation].PauseAtStart()
}

func (a *AnimationData) CancelStandingTimer() {
	if a.StandingTimer != nil {
		a.StandingTimer.Cancel()
		a.StandingTimer = nil
	}
}

var Animation = donburi.NewComponentType[AnimationData]()
