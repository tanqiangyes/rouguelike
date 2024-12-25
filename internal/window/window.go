package window

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/tanqiangyes/rouguelike/internal/config"
	"github.com/tanqiangyes/rouguelike/pkg/logger"
)

// Window 窗口
type Window struct {
	Config *config.Config
}

// NewWindow 创建窗口
func NewWindow(conf *config.Config) *Window {
	return &Window{
		Config: conf,
	}
}

// Run 运行窗口
func (w *Window) Run() {
	ebiten.SetWindowSize(w.Config.Width, w.Config.Height)
	ebiten.SetWindowTitle(w.Config.AppName)
	ebiten.SetWindowIcon([]image.Image{w.Config.Icon})
	ebiten.SetTPS(w.Config.Tps)
	if err := ebiten.RunGame(&Game{}); err != nil {
		logger.NewMainEntry().WithError(err).Error("run game")
	}
}

// Game 游戏
type Game struct{}

// Update 更新
func (g *Game) Update() error {
	return nil
}

// Draw 绘制
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, time.Now().Format(time.StampMilli))
}

// Layout 布局
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
