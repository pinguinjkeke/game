package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *
var assets embed.FS

//go:embed audio/punch.mp3
var Punch_mp3 []byte

//go:embed audio/music.mp3
var Music_mp3 []byte

var ScoreFont = mustLoadFont("font.ttf")

var PlayerSprite = mustLoadImage("player.png")
var CloudsSprite = mustLoadImage("clouds.png")
var MoonSprite = mustLoadImage("moon.png")
var SkySprite = mustLoadImage("sky.png")
var CraneSprite = mustLoadImage("crane.png")
var DrainSprite = mustLoadImage("drain.png")
var RoadSprite = mustLoadImage("road.png")
var SewerSprite = mustLoadImage("sewer.png")
var EnemySprite = mustLoadImage("enemy.png")
var WorldSprite = mustLoadImage("floor.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadFont(name string) font.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
}
