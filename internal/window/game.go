package window

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tanqiangyes/rouguelike/internal/config"
)

// Game 游戏主体
type Game struct {
	Map       *GameMap
	Gd        *GameData
	Config    *config.Config
	World     *ecs.Manager
	WorldTags map[string]ecs.Tag
}

// NewGame 创建游戏主体
func NewGame(conf *config.Config) *Game {
	g := &Game{Gd: NewGameData(), Config: conf}
	g.Map = NewGameMap(g.Gd, conf.Lang)
	world, tags := InitializeWorld(g.Map.CurrentLevel)
	g.WorldTags = tags
	g.World = world
	return g
}

// Update 更新
func (g *Game) Update() error {
	MovePlayer(g)
	return nil
}

// Draw 绘制
func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
}

// Layout 布局
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Gd.TileWidth * g.Gd.ScreenWidth, g.Gd.TileHeight * g.Gd.ScreenHeight
}
