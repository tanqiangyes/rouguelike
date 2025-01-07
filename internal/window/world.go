package window

import (
	"github.com/bytearena/ecs"

	"github.com/tanqiangyes/rouguelike/assets"
)

var positionComponent *ecs.Component
var renderableComponent *ecs.Component

// InitializeWorld 初始化世界
func InitializeWorld(startLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	startRoom := startLevel.Rooms[0]
	x, y := startRoom.Center()
	// 注入主键
	player := manager.NewComponent()
	positionComponent = manager.NewComponent()
	renderableComponent = manager.NewComponent()
	movable := manager.NewComponent()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(positionComponent, &Position{
			X: x,
			Y: y,
		}).
		AddComponent(renderableComponent, &Renderable{
			Image: assets.PlayerImage,
		}).
		AddComponent(movable, Movable{})

	tags["players"] = ecs.BuildTag(player, positionComponent)
	tags["renderables"] = ecs.BuildTag(renderableComponent, positionComponent)
	return manager, tags
}
