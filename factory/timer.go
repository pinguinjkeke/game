package factory

import (
	"game/archetype"
	"game/component"
	"github.com/solarlune/ebitick"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateTimer(ecs *ecs.ECS) *donburi.Entry {
	timerEntry := archetype.Timer.Spawn(ecs)

	component.Timer.Set(timerEntry, ebitick.NewTimerSystem())

	return timerEntry
}
