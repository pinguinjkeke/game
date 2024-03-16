package renderer

import (
	"fmt"
	"game/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func RenderPlayer(ecs *ecs.ECS, screen *ebiten.Image) {
	playerEntry := component.Player.MustFirst(ecs.World)
	player := component.Player.Get(playerEntry)
	animations := component.Animation.Get(playerEntry)
	object := component.Object.Get(playerEntry)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%b %f", player.Jumping, player.SpeedY))

	op := &ganim8.DrawOptions{
		X:       object.Position.X + component.PlayerFrameWidth/4,
		Y:       object.Position.Y + component.PlayerFrameHeight/2,
		ScaleX:  1,
		ScaleY:  1,
		OriginX: 0.5,
		OriginY: 0.5,
	}

	if player.MovingDirection < 0 {
		op.ScaleX = -1
	}

	animations.GetActive().Draw(screen, op)
}
