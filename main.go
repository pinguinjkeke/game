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
	WindowHeight  = 480
	WindowWidth   = 640
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
	ecs.AddSystem(system.UpdatePlayer)
	ecs.AddSystem(system.UpdatePlayerAnimation)

	ecs.AddRenderer(layers.Default, renderer.RenderWall)
	ecs.AddRenderer(layers.Default, renderer.RenderPlayer)

	factory.CreateControls(ecs)
	space := factory.CreateSpace(ecs, WindowWidth, WindowHeight)

	factory.CreateTimer(ecs)

	physics.Add(
		space,
		factory.CreateWall(ecs, 0, 0, WallThickness, WindowHeight),
		factory.CreateWall(ecs, WindowWidth-WallThickness, 0, WallThickness, WindowHeight),
		factory.CreateWall(ecs, 0, 0, WindowWidth, WallThickness),
		factory.CreateWall(ecs, 0, WindowHeight-WallThickness*4, WindowWidth, WallThickness*4),
		factory.CreatePlayer(ecs, WindowWidth/2, WindowHeight/2),
	)

	factory.CreatePlayerAnimation(ecs)

	game := &Game{ecs}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
