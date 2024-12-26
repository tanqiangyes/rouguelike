package window

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/tanqiangyes/rouguelike/internal/components"
)

// ProcessRenderables 渲染
func ProcessRenderables(g *Game, level Level, screen *ebiten.Image) {
	for _, result := range g.World.Query(g.WorldTags["renderables"]) {
		pos := result.Components[components.PositionComponent].(*components.Position)
		img := result.Components[components.RenderableComponent].(*components.Renderable).Image

		index := level.GetIndexFromXY(pos.X, pos.Y)
		tile := level.Tiles[index]
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
		screen.DrawImage(img, op)
	}
}

// MovePlayer 移动玩家
func MovePlayer(g *Game) {
	players := g.WorldTags["players"]

	x := 0
	y := 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
	}

	for _, result := range g.World.Query(players) {
		pos := result.Components[components.PositionComponent].(*components.Position)
		pos.X += x
		pos.Y += y
	}
}
