package system

import (
	"game/component"
	"game/controls"
	"game/physics"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"math"
)

func playerJump(ecs *ecs.ECS, playerEntry *donburi.Entry) {
	spaceEntry := component.Space.MustFirst(ecs.World)
	space := component.Space.Get(spaceEntry)
	player := component.Player.Get(playerEntry)
	playerObject := component.Object.Get(playerEntry)
	playerControls := component.ControlsHandler.Get(playerEntry)

	player.SpeedY += physics.Gravity

	if !player.DoubleJumping && playerControls.ActionIsJustPressed(controls.Jump) && (player.Jumping || player.Ground != nil) {
		player.SpeedY = -physics.JumpSpeed

		if player.Jumping {
			player.DoubleJumping = true
		}

		player.Jumping = true
	}

	player.Ground = nil
	dy := math.Max(math.Min(player.SpeedY, float64(space.Space.CellHeight)), float64(-space.Space.CellHeight))
	checkDistance := dy

	if dy >= 0 {
		checkDistance++
	}

	if collision := playerObject.Check(0, checkDistance, physics.TagSolid); collision != nil {
		if dy < 0 {
			slideOverTopPlatform(playerObject, collision)
		} else {
			contactGround(&dy, player, playerObject, collision)
		}
	}

	if player.Ground != nil {
		player.Jumping = false
		player.DoubleJumping = false
	}

	playerObject.Position.Y += dy
}

func slideOverTopPlatform(playerObject *resolv.Object, collision *resolv.Collision) {
	firstCell := collision.Cells[0]
	slide, sliding := collision.SlideAgainstCell(firstCell, physics.TagSolid)

	if sliding && firstCell.ContainsTags(physics.TagSolid) && math.Abs(slide.X) <= physics.TopPlatformSlideDistance {
		playerObject.Position.X = slide.X
	}
}

func contactGround(
	dy *float64,
	player *component.PlayerData,
	playerObject *resolv.Object,
	collision *resolv.Collision,
) {
	solids := collision.ObjectsByTags(physics.TagSolid)

	if len(solids) == 0 || (player.Ground != nil && player.Ground.Position.Y < solids[0].Position.Y) {
		return
	}

	*dy = collision.ContactWithObject(solids[0]).Y
	player.SpeedY = 0

	if solids[0].Position.Y > playerObject.Position.Y {
		player.Ground = solids[0]
	}
}