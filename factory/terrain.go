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
	startX   int
	startY   int
	vertices []float64
	sprite   *ebiten.Image
}

var objects = [...]object{
	ObjectTypeRv: object{
		startX: 0,
		startY: 0,
		vertices: []float64{
			0, 59,
			116, 59,
			116, 0,
			15, 0,
			9, 26,
			5, 26,
			1, 30,
		},
		sprite: assets.RvSprite,
	},
}

func CreateTerrain(ecs *ecs.ECS, objectType int, x, y float64) *donburi.Entry {
	terrainEntry := archetype.Terrain.Spawn(ecs)

	sprite := objects[objectType].sprite

	object := resolv.NewObject(x, y, float64(sprite.Bounds().Dx()), float64(sprite.Bounds().Dy()), physics.TagSolid)
	object.SetShape(resolv.NewConvexPolygon(0, 0, objects[objectType].vertices...))

	component.Object.Set(terrainEntry, object)
	component.Terrain.Set(terrainEntry, &component.TerrainData{
		Sprite: sprite,
	})

	return terrainEntry
}
