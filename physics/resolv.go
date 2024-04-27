package physics

const (
	TagSolid   = "solid"
	TagTerrain = "terrain"
)

const (
	Friction                 = 0.5
	Acceleration             = 1.0
	RunningAcceleration      = 1.5
	MaxWalkingSpeed          = 1.5
	MaxRunningSpeed          = 4.0
	XSpeedDiff               = MaxRunningSpeed - MaxWalkingSpeed
	LandingDistance          = 280.0
	JumpSpeed                = 10.0
	Gravity                  = 0.75
	TopPlatformSlideDistance = 8
)
