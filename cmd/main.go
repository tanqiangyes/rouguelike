package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Rouguelike")

	w.Resize(fyne.NewSize(800, 600))
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("编辑"),
	),
	)
	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
