package archetype

import (
	"game/component"
	"game/tags"
)

var Sky = newArchetype(
	tags.Background,
	component.Sky,
)
