package factory

import (
	"game/archetype"
	"game/assets"
	"game/component"
	"game/physics"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type ObjectType = int

const (
	ObjectTypeRv ObjectType = iota
)

type object struct {
	startX   float64
	startY   float64
	width    float64
	height   float64
	vertices []float64
	sprite   *ebiten.Image
	tag      string
}

var objects = [...]object{
	ObjectTypeRv: object{
		startX: 18,
		startY: 2,
		width:  102,
		height: 64,
		vertices: []float64{
			0, 1,
			112, 1,
			112, 65,
			0, 65,
		},
		sprite: assets.RvSprite,
		tag:    physics.TagSolid,
	},
}

func CreateTerrain(ecs *ecs.ECS, objectType int, x, y float64) *donburi.Entry {
	terrainEntry := archetype.Terrain.Spawn(ecs)

	object := objects[objectType]
	sprite := object.sprite

	spaceEntry := component.Space.MustFirst(ecs.World)
	space := component.Space.Get(spaceEntry)

	shape := space.Space.AddShape(cp.NewSegment(
		space.Space.StaticBody,
		cp.Vector{X: x, Y: y},
		cp.Vector{X: x + object.width, Y: y + object.height},
		0,
	))
	shape.SetFriction(1)
	shape.SetElasticity(1)

	component.Shape.Set(terrainEntry, &component.ShapeData{
		Shape:  shape,
		Width:  object.width,
		Height: object.height,
	})
	component.Terrain.Set(terrainEntry, &component.TerrainData{
		StartX: object.startX,
		StartY: object.startY,
		Sprite: sprite,
	})

	return terrainEntry
}
