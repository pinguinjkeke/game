package archetype

import (
	"game/component"
	"game/tags"
)

var Wall = newArchetype(
	tags.Wall,
	component.Shape,
)
