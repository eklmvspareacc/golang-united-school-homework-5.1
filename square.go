package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (copy Square) End() Point {
	return Point{x: copy.start.x + int(copy.a), y: copy.start.y + int(copy.a)}
}

func (copy Square) Area() uint {
	return copy.a * copy.a
}

func (copy Square) Perimeter() uint {
	return copy.a * 4
}
