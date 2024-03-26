package system

import (
	"game/component"
	"game/controls"
	"game/physics"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"math"
)

func playerWalk(ecs *ecs.ECS, playerEntry *donburi.Entry) {
	player := component.Player.Get(playerEntry)
	playerObject := component.Object.Get(playerEntry)
	playerControls := component.ControlsHandler.Get(playerEntry)

	acceleration, maxSpeed := physics.Acceleration, physics.MaxRunningSpeed

	player.Running = false
	player.JustStoppedRunning = playerControls.ActionIsJustReleased(controls.Run)
	player.JustChangedMovingDirection = false

	if playerControls.ActionIsPressed(controls.Run) {
		player.Running = true

		acceleration = physics.RunningAcceleration
	}

	movingDirection := 0

	if playerControls.ActionIsPressed(controls.MoveLeft) {
		player.SpeedX -= acceleration
		movingDirection = -1
	}

	if playerControls.ActionIsPressed(controls.MoveRight) {
		player.SpeedX += acceleration
		movingDirection = 1
	}

	if movingDirection != 0 && player.MovingDirection != movingDirection {
		player.MovingDirection = movingDirection
		player.JustLanded = false
		player.JustChangedMovingDirection = true
		player.Running = false
	}

	if player.SpeedX > physics.Friction {
		player.SpeedX -= physics.Friction
	} else if player.SpeedX < -physics.Friction {
		player.SpeedX += physics.Friction
	} else {
		player.SpeedX = 0
	}

	if player.Running {
		if player.Jumping {
			maxSpeed += physics.XSpeedDiff * 2
		}

		if player.JustLanded {
			maxSpeed += physics.XSpeedDiff
		}
	} else if math.Abs(player.SpeedX) > physics.MaxWalkingSpeed {
		maxSpeed = math.Max(math.Abs(player.SpeedX)-physics.XSpeedDiff/4, physics.MaxWalkingSpeed)
	}

	player.SpeedX = math.Max(math.Min(player.SpeedX, maxSpeed), -maxSpeed)

	dx := player.SpeedX

	if check := playerObject.Check(dx, 0, physics.TagSolid, physics.TagTerrain); check != nil {
		dx = check.ContactWithCell(check.Cells[0]).X
		player.SpeedX = 0
	}

	playerObject.Position.X += dx

	if player.JustLanded {
		player.LandDistance += dx
	}
}
