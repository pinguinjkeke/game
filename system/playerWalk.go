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

	if playerControls.ActionIsPressed(controls.Run) {
		player.Running = true
		acceleration = physics.RunningAcceleration
	}

	if playerControls.ActionIsPressed(controls.MoveLeft) {
		player.SpeedX -= acceleration
		player.MovingDirection = -1
	}

	if playerControls.ActionIsPressed(controls.MoveRight) {
		player.SpeedX += acceleration
		player.MovingDirection = 1
	}

	if player.SpeedX > physics.Friction {
		player.SpeedX -= physics.Friction
	} else if player.SpeedX < -physics.Friction {
		player.SpeedX += physics.Friction
	} else {
		player.SpeedX = 0
	}

	if !player.Running && math.Abs(player.SpeedX) > physics.MaxWalkingSpeed {
		maxSpeed = math.Max(math.Abs(player.SpeedX)-(physics.MaxRunningSpeed-physics.MaxWalkingSpeed)/4, physics.MaxWalkingSpeed)
	}

	player.SpeedX = math.Max(math.Min(player.SpeedX, maxSpeed), -maxSpeed)

	dx := player.SpeedX

	if check := playerObject.Check(dx, 0, physics.TagSolid); check != nil {
		dx = check.ContactWithCell(check.Cells[0]).X
		player.SpeedX = 0
	}

	playerObject.Position.X += dx
}
