package archetype

import (
	"game/component"
	"game/tags"
)

var Player = newArchetype(
	tags.Player,
	component.Animation,
	component.Shape,
	component.ControlsHandler,
	component.Player,
)
