package archetype

import (
	"game/component"
	"game/tags"
)

var Space = newArchetype(
	tags.Space,
	component.Space,
)
