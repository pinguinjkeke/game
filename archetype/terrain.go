package archetype

import "game/component"

var Terrain = newArchetype(
	component.Shape,
	component.Terrain,
)
