package window

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// MapTile 地图块
type MapTile struct {
	PixelX     int
	PixelY     int
	Blocked    bool
	Opaque     bool
	IsRevealed bool
	Image      *ebiten.Image
}
