package assets

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed *.png
var iconFile embed.FS

// IconImage app icon
var IconImage *ebiten.Image

// FloorImage app FloorImage
var FloorImage *ebiten.Image

// PlayerImage app PlayerImage
var PlayerImage *ebiten.Image

// WallImage app WallImage
var WallImage *ebiten.Image

// SkellyImage app SkellyImage
var SkellyImage *ebiten.Image

func init() {
	var err error
	IconImage, _, err = ebitenutil.NewImageFromFileSystem(iconFile, "gopher.png")
	if err != nil {
		panic(err)
	}
	FloorImage, _, err = ebitenutil.NewImageFromFileSystem(iconFile, "floor.png")
	if err != nil {
		panic(err)
	}
	PlayerImage, _, err = ebitenutil.NewImageFromFileSystem(iconFile, "player.png")
	if err != nil {
		panic(err)
	}
	WallImage, _, err = ebitenutil.NewImageFromFileSystem(iconFile, "wall.png")
	if err != nil {
		panic(err)
	}
	SkellyImage, _, err = ebitenutil.NewImageFromFileSystem(iconFile, "skelly.png")
	if err != nil {
		panic(err)
	}
}
