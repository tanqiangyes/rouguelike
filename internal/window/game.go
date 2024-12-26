package window

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tanqiangyes/rouguelike/internal/components"
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
	gd := NewGameData()
	world, tags := components.InitializeWorld()
	return &Game{
		Map:       NewGameMap(gd, conf.Lang),
		Gd:        gd,
		Config:    conf,
		World:     world,
		WorldTags: tags,
	}
}

// Update 更新
func (g *Game) Update() error {
	return nil
}

// Draw 绘制
func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.Dungeons[0].Levels[0]
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
}

// Layout 布局
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Gd.TileWidth * g.Gd.ScreenWidth, g.Gd.TileHeight * g.Gd.ScreenHeight
}
