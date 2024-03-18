package main

import (
	"game/factory"
	"game/layers"
	"game/physics"
	"game/renderer"
	"game/system"
	"github.com/yohamta/donburi/ecs"
	"image/color"
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
	screen.Fill(color.RGBA{20, 20, 40, 255})

	g.ecs.Draw(screen)
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
	ecs.AddSystem(system.UpdateBuildings)
	ecs.AddSystem(system.UpdatePlayer)
	ecs.AddSystem(system.UpdatePlayerAnimation)
	ecs.AddSystem(system.UpdateCamera)

	ecs.AddRenderer(layers.CameraClear, renderer.RenderCameraClear)
	ecs.AddRenderer(layers.Background, renderer.RenderBackground)
	ecs.AddRenderer(layers.Game, renderer.RenderWall)
	ecs.AddRenderer(layers.Game, renderer.RenderPlayer)
	ecs.AddRenderer(layers.Camera, renderer.RenderCamera)

	factory.CreateCamera(ecs, WindowWidth, WindowHeight)

	factory.CreateControls(ecs)
	space := factory.CreateSpace(ecs, LevelWidth, LevelHeight)

	factory.CreateTimer(ecs)

	physics.Add(
		space,
		factory.CreateWall(ecs, 0, 0, WallThickness, LevelHeight),
		factory.CreateWall(ecs, LevelWidth-WallThickness, 0, WallThickness, LevelHeight),
		factory.CreateWall(ecs, 0, 0, LevelWidth, WallThickness),
		factory.CreateWall(ecs, 0, LevelHeight-WallThickness*4, LevelWidth, WallThickness*4),
		factory.CreatePlayer(ecs, WindowWidth/2, LevelHeight/2),
	)

	factory.CreatePlayerAnimation(ecs)

	factory.CreateBuildings(ecs, LevelWidth, LevelHeight)
	factory.CreateSky(ecs, LevelWidth, WindowWidth)

	game := &Game{ecs}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
