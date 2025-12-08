package main

import (
	"fmt"
	"os"
	"strings"
)

type Coords struct {
	x, y int
}

func moveTachion(input []string, x, y int, cache map[Coords]int) int {
	_, ok := cache[Coords{x, y}]
	if ok {
		return cache[Coords{x, y}]
	}
	if y >= len(input) {
		cache[Coords{x, y}] = 1
		return 1
	}
	if input[y][x] == '.' {
		val := moveTachion(input, x, y+1, cache)
		cache[Coords{x, y}] = val
		return val
	} else {
		val := moveTachion(input, x-1, y+1, cache) + moveTachion(input, x+1, y+1, cache)
		cache[Coords{x, y}] = val
		return val
	}

}

func main() {

	file, _ := os.ReadFile("day7/input.txt")
	input := strings.Split(string(file), "\n")
	cache := make(map[Coords]int)
	x := strings.Index(input[0], "S")

	fmt.Println(moveTachion(input, x, 1, cache))

}
