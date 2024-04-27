package system

//import (
//	"game/component"
//	"game/controls"
//	"game/physics"
//	"github.com/solarlune/resolv"
//	"github.com/yohamta/donburi"
//	"github.com/yohamta/donburi/ecs"
//	"math"
//)
//
//func playerJump(ecs *ecs.ECS, playerEntry *donburi.Entry) {
//	spaceEntry := component.Space.MustFirst(ecs.World)
//	space := component.Space.Get(spaceEntry)
//	player := component.Player.Get(playerEntry)
//	playerObject := component.Shape.Get(playerEntry)
//	playerControls := component.ControlsHandler.Get(playerEntry)
//
//	player.SpeedY += physics.Gravity
//	player.JustLanded = player.Running && player.JustLanded && math.Abs(player.LandDistance) < physics.LandingDistance
//
//	if !player.Jumping && playerControls.ActionIsJustPressed(controls.Jump) && (player.Jumping || player.Grounded != nil) {
//		player.SpeedY = -physics.JumpSpeed
//		jumpSpeedX := physics.XSpeedDiff
//
//		if player.Running {
//			jumpSpeedX *= 2
//		}
//
//		if player.SpeedX < 0 {
//			jumpSpeedX *= -1
//		}
//
//		player.SpeedX += jumpSpeedX
//
//		player.Jumping = true
//	}
//
//	player.Grounded = nil
//	dy := math.Max(math.Min(player.SpeedY, float64(space.Space.CellHeight)), float64(-space.Space.CellHeight))
//	checkDistance := dy
//
//	if dy >= 0 {
//		checkDistance++
//	}
//
//	if collision := playerObject.Check(0, checkDistance, physics.TagSolid, physics.TagTerrain); collision != nil {
//		if dy < 0 {
//			slideOverTopPlatform(playerObject, collision)
//		} else {
//			contactGround(&dy, player, playerObject, collision)
//		}
//	}
//
//	if player.Grounded != nil {
//		if player.Jumping {
//			player.JustLanded = true
//			player.LandDistance = 0.0
//		}
//
//		player.Jumping = false
//	}
//
//	playerObject.Position.Y += dy
//}
//
//func slideOverTopPlatform(playerObject *resolv.Object, collision *resolv.Collision) {
//	firstCell := collision.Cells[0]
//	slide, sliding := collision.SlideAgainstCell(firstCell, physics.TagSolid)
//
//	if sliding && firstCell.ContainsTags(physics.TagSolid) && math.Abs(slide.X) <= physics.TopPlatformSlideDistance {
//		playerObject.Position.X = slide.X
//	}
//}
//
//func contactGround(
//	dy *float64,
//	player *component.PlayerData,
//	playerObject *resolv.Object,
//	collision *resolv.Collision,
//) {
//	solids := collision.ObjectsByTags(physics.TagSolid)
//
//	if len(solids) != 0 && (player.Grounded == nil || player.Grounded.Position.Y > solids[0].Position.Y) {
//		*dy = collision.ContactWithObject(solids[0]).Y
//		player.SpeedY = 0
//
//		if solids[0].Position.Y > playerObject.Position.Y {
//			player.Grounded = solids[0]
//
//			return
//		}
//	}
//
//	if terrains := collision.ObjectsByTags(physics.TagTerrain); len(terrains) > 0 {
//		if contactSet := playerObject.Shape.Intersection(player.SpeedX, 2, terrains[0].Shape); contactSet != nil {
//			player.Grounded = terrains[0]
//			player.SpeedY = 0
//			*dy = contactSet.TopmostPoint().Y - playerObject.Bottom() + 0.1
//		}
//	}
//}
