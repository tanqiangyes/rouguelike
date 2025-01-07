package window

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tanqiangyes/rouguelike/internal/config"
)

// Game 游戏主体
type Game struct {
	Map         *GameMap
	Gd          *GameData
	Config      *config.Config
	World       *ecs.Manager
	Turn        TurnState
	TurnCounter int
	WorldTags   map[string]ecs.Tag
}

// NewGame 创建游戏主体
func NewGame(conf *config.Config) *Game {
	g := &Game{Gd: NewGameData(), Config: conf}
	g.Map = NewGameMap(g.Gd, conf.Lang)
	world, tags := InitializeWorld(g.Map.CurrentLevel)
	g.WorldTags = tags
	g.World = world
	g.Turn = PlayerTurn
	g.TurnCounter = 0
	return g
}

// Update 更新
func (g *Game) Update() error {
	g.TurnCounter++
	if g.Turn == PlayerTurn && g.TurnCounter > 20 {
		MovePlayer(g)
	}
	g.Turn = PlayerTurn
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
