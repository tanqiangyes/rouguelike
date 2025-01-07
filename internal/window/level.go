package window

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/norendren/go-fov/fov"

	"github.com/tanqiangyes/rouguelike/assets"
	"github.com/tanqiangyes/rouguelike/pkg"
)

// Level  dungeon level.
type Level struct {
	Tiles         []MapTile
	Rooms         []Rect
	PlayerVisible *fov.View
	Gd            *GameData
}

// NewLevel 创建默认的level
func NewLevel(data *GameData) Level {
	l := Level{
		Gd: data,
	}
	// l.Tiles = l.CreateTiles(data)
	l.Rooms = make([]Rect, 0)
	l.GenerateLevelTiles()
	l.PlayerVisible = fov.New()
	return l
}

// GetIndexFromXY 获取坐标对应的索引
func (l *Level) GetIndexFromXY(x int, y int) int {
	return y*l.Gd.ScreenWidth + x
}

// CreateTiles 创建tiles
func (l *Level) CreateTiles(gd *GameData) []MapTile {
	tiles := make([]MapTile, l.Gd.ScreenHeight*l.Gd.ScreenWidth)
	index := 0
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = l.GetIndexFromXY(x, y)
			tiles[index] = MapTile{
				PixelX:     x * gd.TileWidth,
				PixelY:     y * gd.TileHeight,
				Blocked:    true,
				Opaque:     true,
				IsRevealed: false,
				Image:      assets.WallImage,
				TileType:   Wall,
			}
		}
	}
	return tiles
}

// DrawLevel 绘制level
func (l *Level) DrawLevel(screen *ebiten.Image) {
	for x := 0; x < l.Gd.ScreenWidth; x++ {
		for y := 0; y < l.Gd.ScreenHeight; y++ {
			index := l.GetIndexFromXY(x, y)
			tile := l.Tiles[index]
			isVis := l.PlayerVisible.IsVisible(x, y)
			if isVis {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
				screen.DrawImage(tile.Image, op)
				l.Tiles[index].IsRevealed = true
			} else if tile.IsRevealed && !tile.Blocked {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
				op.ColorScale.ScaleWithColor(color.RGBA{
					R: 149,
					G: 165,
					B: 166,
					A: 0,
				})
				// op.ColorScale.Scale(218, 223, 225, 1)
				// op.ColorM.Translate(218, 223, 225, 1)
				screen.DrawImage(tile.Image, op)
			} else if tile.IsRevealed && tile.Blocked {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
				screen.DrawImage(tile.Image, op)
			}
		}
	}
}

// CreateRoom 创建房间
func (l *Level) CreateRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := l.GetIndexFromXY(x, y)
			l.Tiles[index].Blocked = false
			l.Tiles[index].Opaque = false
			l.Tiles[index].IsRevealed = false
			l.Tiles[index].TileType = Floor
			l.Tiles[index].Image = assets.FloorImage
		}
	}
}

// GenerateLevelTiles 创建层级tile
func (l *Level) GenerateLevelTiles() {
	containsRooms := false
	MinSize := 6
	MaxSize := 10
	MaxRooms := 30

	gd := l.Gd
	tiles := l.CreateTiles(gd)
	l.Tiles = tiles

	for idx := 0; idx < MaxRooms; idx++ {
		w := pkg.GetRandomBetween(MinSize, MaxSize)
		h := pkg.GetRandomBetween(MinSize, MaxSize)
		x := pkg.GetDiceRoll(gd.ScreenWidth - w - 1)
		y := pkg.GetDiceRoll(gd.ScreenHeight - h - 1)

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
			if containsRooms {
				newX, newY := newRoom.Center()
				prevX, prevY := l.Rooms[len(l.Rooms)-1].Center()

				coinflip := pkg.GetDiceRoll(2)

				if coinflip == 2 {
					l.createHorizontalTunnel(prevX, newX, prevY)
					l.createVerticalTunnel(prevY, newY, newX)
				} else {
					l.createHorizontalTunnel(prevX, newX, newY)
					l.createVerticalTunnel(prevY, newY, prevX)
				}
			}
			l.Rooms = append(l.Rooms, newRoom)
			containsRooms = true
		}
	}
}

func (l *Level) createHorizontalTunnel(x1 int, x2 int, y int) {
	x := pkg.Min(x1, x2)
	max := pkg.Max(x1, x2) + 1
	for ; x < max; x++ {
		index := l.GetIndexFromXY(x, y)
		if index > 0 && index < l.Gd.ScreenWidth*l.Gd.ScreenHeight {
			l.Tiles[index].Blocked = false
			l.Tiles[index].Image = assets.FloorImage
			l.Tiles[index].TileType = Floor
		}
	}
}
func (l *Level) createVerticalTunnel(y1 int, y2 int, x int) {
	y := pkg.Min(y1, y2)
	max := pkg.Max(y1, y2) + 1
	for ; y < max; y++ {
		index := l.GetIndexFromXY(x, y)
		if index > 0 && index < l.Gd.ScreenWidth*l.Gd.ScreenHeight {
			l.Tiles[index].Blocked = false
			l.Tiles[index].Image = assets.FloorImage
			l.Tiles[index].TileType = Floor
		}
	}
}

// InBounds 是否在棋牌内
func (l Level) InBounds(x, y int) bool {
	if x < 0 || x > l.Gd.ScreenWidth || y < 0 || y > l.Gd.ScreenHeight {
		return false
	}
	return true
}

// IsOpaque 是否可见
func (l Level) IsOpaque(x, y int) bool {
	idx := l.GetIndexFromXY(x, y)
	return l.Tiles[idx].TileType.IsWall()
}
