package window

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	"github.com/tanqiangyes/rouguelike/internal/config"
)

// Window 窗口
type Window struct {
	Config     config.Config
	MainWindow fyne.Window
	fyne.App
}

// NewWindow 创建窗口
func NewWindow(conf config.Config) *Window {
	return &Window{
		Config: conf,
		App:    app.New(),
	}
}
func (w *Window) Create() {
	window := w.NewWindow("Rouguelike")
	w.MainWindow = window

	window.Resize(fyne.NewSize(800, 600))
	window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("编辑"),
		Resize(window),
	),
	)
	clock := widget.NewLabel("")
	updateTime(clock)

	window.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
}

// Run 运行窗口
func (w *Window) Run() {
	w.Create()
	w.MainWindow.Show()
	w.App.Run()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
func Resize(w fyne.Window) *fyne.Menu {
	resize := fyne.NewMenu("修改大小")
	resize.Items = append(resize.Items, fyne.NewMenuItem("800x600", func() {
		w.Resize(fyne.NewSize(800, 600))
	}), fyne.NewMenuItem("1000x1000", func() {
		w.Resize(fyne.NewSize(1000, 1000))
	}),
	)
	return resize
}
