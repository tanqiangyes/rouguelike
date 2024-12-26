package window

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tanqiangyes/rouguelike/assets"
	"github.com/tanqiangyes/rouguelike/internal/config"
	"github.com/tanqiangyes/rouguelike/pkg/logger"
)

// Window 游戏窗口
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
	ebiten.SetWindowIcon([]image.Image{assets.IconImage})
	ebiten.SetTPS(w.Config.Tps)
	if err := ebiten.RunGame(NewGame(w.Config)); err != nil {
		logger.NewMainEntry().WithError(err).Error("run game")
	}
}
