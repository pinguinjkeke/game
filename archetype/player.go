package archetype

import (
	"game/component"
	"game/tags"
)

var Player = newArchetype(
	tags.Player,
	component.Animation,
	component.Object,
	component.ControlsHandler,
	component.Player,
)
