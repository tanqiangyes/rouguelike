package window

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tanqiangyes/rouguelike/assets"
	"github.com/tanqiangyes/rouguelike/pkg"
)

// Level  dungeon level.
type Level struct {
	Tiles []MapTile
	Rooms []Rect
	Gd    *GameData
}

// NewLevel 创建默认的level
func NewLevel(data *GameData) Level {
	l := Level{
		Gd: data,
	}
	l.Tiles = l.CreateTiles(data)
	l.Rooms = make([]Rect, 0)
	l.GenerateLevelTiles()
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
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = l.GetIndexFromXY(x, y)
			tiles[index] = MapTile{
				PixelX:  x * gd.TileWidth,
				PixelY:  y * gd.TileHeight,
				Blocked: true,
				Opaque:  true,
				Image:   assets.WallImage,
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

// CreateRoom 创建房间
func (l *Level) CreateRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := l.GetIndexFromXY(x, y)
			l.Tiles[index].Blocked = false
			l.Tiles[index].Image = assets.FloorImage
		}
	}
}

// GenerateLevelTiles 创建层级tile
func (l *Level) GenerateLevelTiles() {
	MinSize := 6
	MaxSize := 10
	MaxRooms := 30

	gd := l.Gd
	tiles := l.CreateTiles(gd)
	l.Tiles = tiles

	for idx := 0; idx < MaxRooms; idx++ {
		w := pkg.GetRandomBetween(MinSize, MaxSize)
		h := pkg.GetRandomBetween(MinSize, MaxSize)
		x := pkg.GetDiceRoll(gd.ScreenWidth-w-1) - 1
		y := pkg.GetDiceRoll(gd.ScreenHeight-h-1) - 1

		newRoom := NewRect(x, y, w, h)

		okToAdd := true

		for _, otherRoom := range l.Rooms {
			if newRoom.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}

		if okToAdd {
			l.CreateRoom(newRoom)
			l.Rooms = append(l.Rooms, newRoom)
		}
	}
}
