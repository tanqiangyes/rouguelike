package window

// TileType 类型
type TileType int

// TileType 类型
const (
	Wall TileType = iota
	Floor
)

// IsWall 是否是墙壁
func (tt TileType) IsWall() bool {
	return tt == Wall
}
