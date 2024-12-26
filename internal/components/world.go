package components

import (
	"github.com/bytearena/ecs"
)

func InitializeWorld() (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	// 注入主键
	player := manager.NewComponent()
	position := manager.NewComponent()
	renderable := manager.NewComponent()
	movable := manager.NewComponent()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(position, &Position{
			X: 40,
			Y: 25,
		}).
		AddComponent(renderable, Renderable{}).
		AddComponent(movable, Movable{})

	tags["players"] = ecs.BuildTag(player, position)
	return manager, tags
}
