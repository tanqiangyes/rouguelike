package window

// Rect rect
type Rect struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

// NewRect 新建Rect
func NewRect(x int, y int, width int, height int) Rect {
	return Rect{
		X1: x,
		X2: y,
		Y1: width,
		Y2: height,
	}
}

// Center 中心点
func (r *Rect) Center() (int, int) {
	return (r.X1 + r.X2) / 2, (r.Y1 + r.Y2) / 2
}

// Intersect 判断是否相交
func (r *Rect) Intersect(other Rect) bool {
	return r.X1 <= other.X2 && r.X2 >= other.X1 && r.Y1 <= other.Y1 && r.Y2 >= other.Y2
}
