package window

// GameData 游戏相应的数据
type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

// NewGameData 创建游戏数据
func NewGameData() *GameData {
	return &GameData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TileWidth:    16,
		TileHeight:   16,
	}
}
