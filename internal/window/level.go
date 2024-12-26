package window

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tanqiangyes/rouguelike/assets"
)

// Level  dungeon level.
type Level struct {
	Tiles []MapTile
	Gd    *GameData
}

// NewLevel 创建默认的level
func NewLevel(data *GameData) Level {
	l := Level{
		Gd: data,
	}
	l.Tiles = l.CreateTiles(data)
	return l
}

// GetIndexFromXY 获取坐标对应的索引
func (l *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return y*gd.ScreenWidth + x
}

// CreateTiles 创建tiles
func (l *Level) CreateTiles(gd *GameData) []MapTile {
	tiles := make([]MapTile, l.Gd.ScreenHeight*l.Gd.ScreenWidth)
	index := 0
	gdScreenWidth := gd.ScreenWidth - 1
	gdScreenHeight := gd.ScreenHeight - 1
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = l.GetIndexFromXY(x, y)
			if x == 0 || x == gdScreenWidth || y == 0 || y == gdScreenHeight {
				tiles[index] = MapTile{
					PixelX:  x * gd.TileWidth,
					PixelY:  y * gd.TileHeight,
					Blocked: true,
					Opaque:  true,
					Image:   assets.WallImage,
				}
			} else {
				tiles[index] = MapTile{
					PixelX:  x * gd.TileWidth,
					PixelY:  y * gd.TileHeight,
					Blocked: false,
					Opaque:  false,
					Image:   assets.FloorImage,
				}
			}
		}
	}
	return tiles
}

// DrawLevel 绘制level
func (l *Level) DrawLevel(screen *ebiten.Image) {
	for x := 0; x < l.Gd.ScreenWidth; x++ {
		for y := 0; y < l.Gd.ScreenHeight; y++ {
			tile := l.Tiles[l.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}
