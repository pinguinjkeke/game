package main

import (
	"game/component"
	"game/factory"
	"game/layers"
	"game/renderer"
	"game/system"
	"github.com/yohamta/donburi/ecs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
)

const (
	WindowWidth   = 640
	WindowHeight  = 480
	LevelHeight   = WindowHeight
	LevelWidth    = WindowWidth * 10
	WallThickness = 4
)

type Game struct {
	ecs *ecs.ECS
}

func (g *Game) Update() error {
	g.ecs.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	cameraEntry := component.Camera.MustFirst(g.ecs.World)
	camera := component.Camera.Get(cameraEntry)

	camera.Surface.Clear()
	g.ecs.Draw(camera.Surface)
	camera.Blit(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WindowWidth, WindowHeight
}

func main() {
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hello, World!")

	world := donburi.NewWorld()
	ecs := ecs.NewECS(world)

	ecs.AddSystem(system.UpdateTimer)
	ecs.AddSystem(system.UpdateSpace)
	ecs.AddSystem(system.UpdatePlayer)
	ecs.AddSystem(system.UpdatePlayerAnimation)
	ecs.AddSystem(system.UpdateCamera)
	ecs.AddSystem(system.UpdateSky)

	ecs.AddRenderer(layers.Background, renderer.RenderBackground)
	ecs.AddRenderer(layers.Background, renderer.RenderRoad)
	ecs.AddRenderer(layers.Game, renderer.RenderPlayer)
	ecs.AddRenderer(layers.Game, renderer.RenderTerrain)

	factory.CreateCamera(ecs, WindowWidth, WindowHeight)

	factory.CreateControls(ecs)
	factory.CreateTimer(ecs)

	factory.CreateSpace(ecs, LevelWidth, LevelHeight)
	factory.CreateWall(ecs, 0, 0, 0, LevelHeight)
	factory.CreateWall(ecs, LevelWidth, 0, 0, LevelHeight)
	factory.CreateWall(ecs, 0, 0, LevelWidth, 0)
	factory.CreateWall(ecs, 0, LevelHeight, LevelWidth, 0)
	factory.CreatePlayer(ecs, WindowWidth/2, LevelHeight/2)
	//factory.CreateTerrain(ecs, factory.ObjectTypeRv, float64(WindowWidth)-100, 400)

	factory.CreatePlayerAnimation(ecs)

	factory.CreateBuildings(ecs)
	factory.CreateSky(ecs)
	factory.CreateRoad(ecs)

	game := &Game{ecs}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
