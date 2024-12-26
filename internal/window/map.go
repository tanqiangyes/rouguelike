package window

// GameMap 地图中的地下城
type GameMap struct {
	Dungeons []Dungeon
}

// NewGameMap 创建一个新的地图
func NewGameMap(data *GameData) *GameMap {
	l := NewLevel(data)
	levels := make([]Level, 0)
	levels = append(levels, l)
	d := Dungeon{
		Name:   "Default",
		Levels: levels,
	}
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, d)
	return &GameMap{Dungeons: dungeons}
}
