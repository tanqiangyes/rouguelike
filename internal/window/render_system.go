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
