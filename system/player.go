package system

import (
	"game/component"
	"github.com/yohamta/donburi/ecs"
)

func UpdatePlayer(ecs *ecs.ECS) {
	playerEntry := component.Player.MustFirst(ecs.World)

	//playerJump(ecs, playerEntry)
	playerWalk(ecs, playerEntry)
}
