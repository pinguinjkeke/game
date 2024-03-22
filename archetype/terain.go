package archetype

import "game/component"

var Terrain = newArchetype(
	component.Object,
	component.Terrain,
)
