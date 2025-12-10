package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func Area(a, b Point) int {
	x := a.x - b.x
	y := a.y - b.y
	val := (max(-y, y) + 1) * (max(-x, x) + 1)
	return max(val, -val)
}

func main() {
	file, _ := os.ReadFile("day9/test.txt")

	lines := strings.Split(string(file), "\n")

	input := make([]Point, len(lines))
	for i, line := range lines {
		val := Point{}
		_, _ = fmt.Sscanf(line, "%d,%d", &val.x, &val.y)
		fmt.Println(val)
		input[i] = val
	}
	maxArea := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			area := Area(input[i], input[j])
			maxArea = max(maxArea, area)
		}
	}

	fmt.Printf("%v\n", maxArea)
}
