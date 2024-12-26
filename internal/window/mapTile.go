package window

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// MapTile represents a single tile on the map.
type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Opaque  bool
	Image   *ebiten.Image
}
