package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Vec3 struct {
	x, y, z     int
	connections []*Vec3
}

func sqDistance(a, b *Vec3) int {
	x := a.x - b.x
	y := a.y - b.y
	z := a.z - b.z
	val := x*x + y*y + z*z
	return val

}
func visit(a *Vec3, visited *[]*Vec3) {
	if a == nil {
		return
	}
	*visited = append(*visited, a)
	for _, child := range a.connections {
		if slices.Index(*visited, child) >= 0 {
			continue
		}
		visit(child, visited)
	}
}

func main() {

	connections := 10000

	file, _ := os.ReadFile("day8/input.txt")
	input := strings.Split(string(file), "\n")

	values := make([]*Vec3, len(input))

	for i := 0; i < len(input); i++ {
		strs := strings.Split(input[i], ",")
		v1, _ := strconv.ParseInt(strs[0], 10, 64)
		v2, _ := strconv.ParseInt(strs[1], 10, 64)
		v3, _ := strconv.ParseInt(strs[2], 10, 64)
		values[i] = &Vec3{
			int(v1),
			int(v2),
			int(v3),
			[]*Vec3{},
		}
	}
	var last1, last2 *Vec3
	for w := 0; w < connections; w++ {
		visited := make([]*Vec3, 0)
		visit(values[0], &visited)
		if len(visited) == len(values) {
			fmt.Printf("count: %v \n", w)
			break
		}

		closest := make(map[*Vec3]*Vec3)

		for i := 0; i < len(values); i++ {
			curval := values[i]
			mindist := math.MaxInt
			curclosest := -1

			if _, ok := closest[curval]; ok {
				continue
			}

			for j := 0; j < len(values); j++ {
				if i == j {
					continue
				}
				if slices.Index(values[j].connections, curval) >= 0 {
					continue
				}
				if _, ok := closest[values[j]]; ok {
					continue
				}
				if curdist := sqDistance(curval, values[j]); curdist <= mindist {
					mindist = curdist
					curclosest = j
				}
			}
			if curclosest == -1 {
				continue
			}
			closest[curval] = values[curclosest]
		}
		mindist := math.MaxInt
		var curmin *Vec3
		for from, to := range closest {
			if dist := sqDistance(from, to); dist < mindist || curmin == nil {

				mindist = dist
				curmin = from
			}
		}
		curmin.connections = append(curmin.connections, closest[curmin])
		closest[curmin].connections = append(closest[curmin].connections, curmin)

		last1 = curmin
		last2 = closest[curmin]

	}

	visited := make([]*Vec3, 0)
	networks := 0
	var sizes []int
	for _, val := range values {
		if slices.Index(visited, val) < 0 {
			last := len(visited)
			visit(val, &visited)
			sizes = append(sizes, len(visited)-last)
			networks += 1
		}
	}
	slices.SortFunc(sizes, func(a, b int) int {
		return b - a
	})
	fmt.Println(networks, sizes)
	fmt.Printf("%v\n", last1.x*last2.x)

}
