package sol

type Point struct {
	x, y int
}
type DetectSquares struct {
	hash map[Point]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		hash: make(map[Point]int),
	}
}

func (this *DetectSquares) Add(point []int) {
	p := Point{x: point[0], y: point[1]}
	(*this).hash[p] += 1
}

func (this *DetectSquares) Count(point []int) int {
	var abs = func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	p := Point{x: point[0], y: point[1]}
	// detect all possible diagonal
	res := 0
	for coord, count := range (*this).hash {
		if abs(coord.x-p.x) != abs(coord.y-p.y) || (coord.x == p.x && coord.y == p.y) {
			continue
		}
		res += count * (*this).hash[Point{x: coord.x, y: p.y}] * (*this).hash[Point{x: p.x, y: coord.y}]
	}
	return res
}
