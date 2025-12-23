package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var connections [][]int

func countExit(node int) int {
	if node < 0 {
		return 1
	}
	res := 0
	for _, con := range connections[node] {
		res += countExit(con)
	}
	return res
}

func main() {
	file, _ := os.ReadFile("day11/input.txt")
	you, out := -1, -1
	lines := strings.Split(string(file), "\n")
	names := make([]string, len(lines))
	connections = make([][]int, len(lines))
	for i, line := range lines {
		if line[:3] == "you" {
			you = i
		}
		names[i] = line[:3]
	}
	for i, line := range lines {
		for _, conn := range strings.Split(line[5:], " ") {
			connections[i] = append(connections[i], slices.Index(names, conn))
		}

	}
	fmt.Println(names)
	fmt.Println(you, out)
	fmt.Println(connections)
	fmt.Println(countExit(you))
}
