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

var CloudsSprite = mustLoadImage("background/clouds.png")
var MoonSprite = mustLoadImage("background/moon.png")
var SkySprite = mustLoadImage("background/sky.png")
var CraneSprite = mustLoadImage("background/crane.png")

var BustStopSprite = mustLoadImage("objects/bus_stop.png")
var RvSprite = mustLoadImage("objects/rv.png")

var DrainSprite = mustLoadImage("road/drain.png")
var RoadSprite = mustLoadImage("road/road.png")
var SewerSprite = mustLoadImage("road/sewer.png")

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
