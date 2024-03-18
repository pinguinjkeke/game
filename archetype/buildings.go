package archetype

import (
	"game/component"
	"game/tags"
)

var Buildings = newArchetype(
	tags.Background,
	component.Buildings,
)
