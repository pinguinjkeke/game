package physics

import (
	"game/chipmunk"
	"game/component"
	"game/controls"
	"github.com/jakecoffman/cp/v2"
	"github.com/yohamta/donburi"
)

func PlayerUpdateVelocity(playerEntry *donburi.Entry) func(body *cp.Body, gravity cp.Vector, damping float64, dt float64) {
	return func(body *cp.Body, gravity cp.Vector, damping, dt float64) {
		player := component.Player.Get(playerEntry)
		playerControls := component.ControlsHandler.Get(playerEntry)
		shape := component.Shape.Get(playerEntry)

		velocity, _ := chipmunk.WalkVelocity, chipmunk.MaxHorizontalVelocity

		player.Running = false
		player.JustStoppedRunning = playerControls.ActionIsJustReleased(controls.Run)
		player.JustChangedMovingDirection = false

		if playerControls.ActionIsPressed(controls.Run) {
			player.Running = true

			velocity = chipmunk.RunVelocity
		}

		// Grab the grounding normal from last frame
		groundNormal := cp.Vector{}
		body.EachArbiter(func(arbiter *cp.Arbiter) {
			if n := arbiter.Normal().Neg(); n.Y < groundNormal.Y {
				groundNormal = n
			}
		})

		grounded := groundNormal.Y < 0

		if player.Grounded != grounded {
			player.Grounded = grounded

			if grounded {
				player.Land(body.Position().X)
			}
		}

		if player.Grounded {
			player.Jumping = false
			player.FinishLanding(body.Position().X)
		}

		body.UpdateVelocity(gravity, damping, dt)

		// Target horizontal speed for air/ground control
		var targetVx, targetVy float64
		var movingDirection int

		if playerControls.ActionIsPressed(controls.MoveLeft) {
			targetVx -= velocity
			movingDirection = -1
		}

		if playerControls.ActionIsPressed(controls.MoveRight) {
			targetVx += velocity
			movingDirection = 1
		}

		if movingDirection != 0 && player.MovingDirection != movingDirection {
			player.MovingDirection = movingDirection
			player.JustLanded = false
			player.LandDistance = 0
			player.JustChangedMovingDirection = true
			player.Running = false
		}

		if !player.Jumping && playerControls.ActionIsJustPressed(controls.Jump) {
			player.Jumping = true
			player.JustLanded = false
			player.LandDistance = 0
			player.Grounded = false
			body.SetVelocityVector(body.Velocity().Add(cp.Vector{Y: -chipmunk.JumpVelocity}))
		}

		// Update the surface velocity and friction
		// Note that the "feet" move in the opposite direction of the player.
		surfaceV := cp.Vector{X: -targetVx}
		shape.SetSurfaceV(surfaceV)
		if player.Grounded {
			shape.SetFriction(chipmunk.PlayerGroundAcceleration / chipmunk.Gravity)
		} else {
			shape.SetFriction(0)
		}

		v := body.Velocity()

		// Apply air control if not grounded
		if !player.Grounded {
			targetVx = cp.LerpConst(v.X, targetVx, chipmunk.PlayerAirAcceleration*dt)
			targetVy = cp.Clamp(v.Y, -chipmunk.FallVelocity, cp.INFINITY)
		} else {
			targetVy = 0
		}

		body.SetVelocity(targetVx, targetVy)
	}
}
