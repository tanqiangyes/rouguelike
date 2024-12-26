package window

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game 游戏实体
type Game struct {
	Map *GameMap
	Gd  *GameData
}

// NewGame 创建游戏实体
func NewGame() *Game {
	gd := NewGameData()
	return &Game{
		Map: NewGameMap(gd),
		Gd:  gd,
	}
}

// Update 更新
func (g *Game) Update() error {
	return nil
}

// Draw 绘制
func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.Dungeons[0].Levels[0]
	// Draw the Map
	for x := 0; x < g.Gd.ScreenWidth; x++ {
		for y := 0; y < g.Gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

// Layout 布局
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Gd.TileWidth * g.Gd.ScreenWidth, g.Gd.TileHeight * g.Gd.ScreenHeight
}
