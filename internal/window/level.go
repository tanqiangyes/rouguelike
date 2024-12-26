package window

import (
	"github.com/tanqiangyes/rouguelike/assets"
)

// Level holds the tile information for a complete dungeon level.
type Level struct {
	Tiles []MapTile
}

// NewLevel creates a new level with the default tiles.
func NewLevel(data *GameData) Level {
	l := Level{}
	l.Tiles = l.CreateTiles(data)
	return l
}

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (l *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return y*gd.ScreenWidth + x
}

// CreateTiles creates a slice of MapTile structs for the game map.
func (l *Level) CreateTiles(gd *GameData) []MapTile {
	tiles := make([]MapTile, 0)
	gdScreenWidth := gd.ScreenWidth - 1
	gdScreenHeight := gd.ScreenHeight - 1
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			if x == 0 || x == gdScreenWidth || y == 0 || y == gdScreenHeight {
				tiles = append(tiles, MapTile{
					PixelX:  x * gd.TileWidth,
					PixelY:  y * gd.TileHeight,
					Blocked: true,
					Opaque:  true,
					Image:   assets.WallImage,
				})
			} else {
				tiles = append(tiles, MapTile{
					PixelX:  x * gd.TileWidth,
					PixelY:  y * gd.TileHeight,
					Blocked: false,
					Opaque:  false,
					Image:   assets.FloorImage,
				})
			}
		}
	}
	return tiles
}
