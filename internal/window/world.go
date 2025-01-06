package window

import (
	"github.com/bytearena/ecs"

	"github.com/tanqiangyes/rouguelike/assets"
)

var PositionComponent *ecs.Component
var RenderableComponent *ecs.Component

// InitializeWorld 初始化世界
func InitializeWorld(startLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	startRoom := startLevel.Rooms[0]
	x, y := startRoom.Center()
	// 注入主键
	player := manager.NewComponent()
	PositionComponent = manager.NewComponent()
	RenderableComponent = manager.NewComponent()
	movable := manager.NewComponent()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(PositionComponent, &Position{
			X: x,
			Y: y,
		}).
		AddComponent(RenderableComponent, &Renderable{
			Image: assets.PlayerImage,
		}).
		AddComponent(movable, Movable{})

	tags["players"] = ecs.BuildTag(player, PositionComponent)
	tags["renderables"] = ecs.BuildTag(RenderableComponent, PositionComponent)
	return manager, tags
}
