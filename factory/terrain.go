package factory

import (
	"game/archetype"
	"game/assets"
	"game/component"
	"game/physics"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
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

	resolvObject := resolv.NewObject(x, y, object.width, object.height, object.tag)
	resolvObject.SetShape(resolv.NewConvexPolygon(x, y, object.vertices...))

	component.Object.Set(terrainEntry, resolvObject)
	component.Terrain.Set(terrainEntry, &component.TerrainData{
		StartX: object.startX,
		StartY: object.startY,
		Sprite: sprite,
	})

	return terrainEntry
}
