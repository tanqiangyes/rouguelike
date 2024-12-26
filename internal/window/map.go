package window

import (
	"github.com/tanqiangyes/rouguelike/pkg/i18n"
)

// GameMap 地图
type GameMap struct {
	Dungeons []Dungeon
	Lang     i18n.Lang
}

// NewGameMap 创建一个新的地图
func NewGameMap(data *GameData, lang i18n.Lang) *GameMap {
	l := NewLevel(data)
	levels := make([]Level, 0)
	levels = append(levels, l)
	d := Dungeon{
		Name:   i18n.Translate("default_dungeon_name", lang),
		Levels: levels,
	}
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, d)
	return &GameMap{
		Dungeons: dungeons,
		Lang:     lang,
	}
}
