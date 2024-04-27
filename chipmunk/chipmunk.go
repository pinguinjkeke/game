package chipmunk

const (
	Gravity                      = 2000.0
	WalkVelocity                 = 150.0
	RunVelocity                  = 280.0
	XVelocityDiff                = RunVelocity - WalkVelocity
	MaxHorizontalVelocity        = 1500.0
	JumpVelocity                 = 500.0
	FallVelocity                 = 900.0
	LandingDistance              = 280.0
	PlayerGroundAccelerationTime = 0.1
	PlayerGroundAcceleration     = WalkVelocity / PlayerGroundAccelerationTime
	PlayerAirAccelerationTime    = 0.25
	PlayerAirAcceleration        = WalkVelocity / PlayerAirAccelerationTime
)
