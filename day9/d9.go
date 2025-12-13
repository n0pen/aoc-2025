package main

import (
	"cmp"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

// Edge is always aligned to an axis
type Edge struct {
	a, b Point
}

/*
Crossing defined as:

	self goes over any point in e
*/

func between[T cmp.Ordered](val, a, b T) bool {
	return val > min(a, b) && val < max(a, b)
}
func betweenInclusive[T cmp.Ordered](val, a, b T) bool {
	return val >= min(a, b) && val <= max(a, b)
}

func (e1 Edge) isCrossing(e2 Edge) bool {
	if e1.isParallelTo(e2) {
		return false
	}

	if !e1.isXAligned() {
		xbetween := between(e1.a.x, e2.a.x, e2.b.x)
		ybetween := between(e2.a.y, e1.a.y, e1.b.y)
		return xbetween &&
			ybetween
	}

	return between(e2.a.x, e1.a.x, e1.b.x) &&
		between(e1.a.y, e2.a.y, e2.b.y)
}

func (e1 Edge) isParallelTo(e2 Edge) bool {
	return e1.isXAligned() == e2.isXAligned()
}

func (e1 Edge) isCollinearTo(e2 Edge) bool {
	if !e1.isParallelTo(e2) {
		return false
	}
	if e1.isXAligned() {
		return e1.a.y == e2.a.y
	}
	return e1.a.x == e2.b.x
}

func (e1 Edge) isXAligned() bool {
	return e1.a.y == e1.b.y
}

func Area(a, b Point) int {
	x := a.x - b.x
	y := a.y - b.y
	val := (max(-y, y) + 1) * (max(-x, x) + 1)
	return max(val, -val)
}

type Polygon []Point

func (pol1 Polygon) crossesPolygon(p2 Polygon) bool {

	colinears := 0
	for j := 0; j < len(pol1); j++ {
		pa, pb := pol1[j], pol1[(j+1)%len(pol1)]
		pe1 := Edge{pa, pb}
		collinear := false
		for i := 0; i < len(p2); i++ {
			pa, pb := p2[i], p2[(i+1)%len(p2)]
			pedge := Edge{pa, pb}
			if pedge.isCrossing(pe1) {
				return true
			}
			if pe1.a == pa && pe1.b == pb || pe1.a == pb && pe1.b == pa {
				collinear = true
			}
			if pe1.isCollinearTo(pedge) {

				if pe1.isXAligned() {

					if betweenInclusive(pedge.a.x, pe1.a.x, pe1.b.x) {
						collinear = between(pedge.a.x, pe1.a.x, pe1.b.x)
						poly := pe1.a.y - pol1[(2+j)%len(pol1)].y
						pedgey := pedge.a.y - p2[(len(p2)-1+i)%len(p2)].y

						if pedge.a.x == pe1.a.x || pedge.a.x == pe1.b.x {

						} else if poly*pedgey > 0 {
							return true
						}

					}
					if betweenInclusive(pedge.b.x, pe1.a.x, pe1.b.x) {
						collinear = between(pedge.b.x, pe1.a.x, pe1.b.x)
						poly := pe1.a.y - pol1[(2+j)%len(pol1)].y
						pedgey := pedge.b.y - p2[(i+2)%len(p2)].y

						if pedge.b.x == pe1.a.x || pedge.b.x == pe1.b.x {

						} else if poly*pedgey > 0 {
							return true
						}

					}
				} else {
					if betweenInclusive(pedge.a.y, pe1.a.y, pe1.b.y) {
						collinear = true
						polx := pe1.a.x - pol1[(2+j)%len(pol1)].x
						pedgex := p2[i].x - p2[(len(p2)-1+i)%len(p2)].x

						if pedge.a.y == pe1.a.y || pedge.a.y == pe1.b.y {

						} else if polx*pedgex > 0 {
							return true
						}

					}
					if betweenInclusive(pedge.b.y, pe1.a.y, pe1.b.y) {
						collinear = true
						polx := pe1.a.x - pol1[(2+j)%len(pol1)].x
						pedgex := pedge.b.x - p2[(i+2)%len(p2)].x
						if pedge.b.y == pe1.a.y || pedge.b.y == pe1.b.y {

						} else if polx*pedgex > 0 {
							return true
						}

					}
				}
			}
		}
		if collinear {
			colinears += 1
		}
	}

	return false

}

func main() {
	file, _ := os.ReadFile("day9/input.txt")

	lines := strings.Split(string(file), "\n")

	input := make(Polygon, len(lines))
	for i, line := range lines {
		val := Point{}
		_, _ = fmt.Sscanf(line, "%d,%d", &val.x, &val.y)
		input[i] = val
	}
	maxArea := 0

	points := ""

	var maxx, maxy int
	var p1, p2 Point
	for i := 0; i < len(input); i++ {

		for j := i + 1; j < len(input); j++ {
			square := Polygon{
				input[i],
				Point{input[i].x, input[j].y},
				input[j],
				Point{input[j].x, input[i].y},
			}
			if square.crossesPolygon(input) {
				continue
			}
			area := Area(input[i], input[j])
			if area > maxArea {
				maxArea = area
				p1 = input[i]
				p2 = input[j]
			}

		}
		maxx = max(maxx, input[i].x)
		maxy = max(maxy, input[i].y)
		points += fmt.Sprintf("%v,%v ", input[i].x, input[i].y)
	}

	square := fmt.Sprintf("%d,%d %d,%d %d,%d %d,%d", p1.x, p1.y, p1.x, p2.y, p2.x, p2.y, p2.x, p1.y)
	data := fmt.Sprintf(`
			<svg height="200" width="200" viewBox="0 0 %d %d"  xmlns="http://www.w3.org/2000/svg">
				<polygon points="%s" style="fill:lime;"/>
				<polygon points="%s" style="stroke:#5556;stroke-width:%f;fill:transparent"/>
			</svg>
		`, maxx+maxx/10, maxy+maxy/10, points, square, float64(maxy)/100)

	_ = os.WriteFile("day9/d9.svg", []byte(data), 0666)

	fmt.Printf("%v\n", maxArea)
}
