package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var connections [][]int

func countExit(node, target int, cache map[int]int) int {
	if cache == nil {
		cache = make(map[int]int)
	}
	if val, res := cache[node]; res {
		return val
	}
	if node == target {
		return 1
	}
	if node == -1 {
		return 0
	}
	res := 0
	for _, con := range connections[node] {
		res += countExit(con, target, cache)
	}

	cache[node] = res

	return res
}

var you, svr, fft, dac int

func main() {
	file, _ := os.ReadFile("day11/input.txt")
	out := -1

	lines := strings.Split(string(file), "\n")
	names := make([]string, len(lines))

	connections = make([][]int, len(lines))
	for i, line := range lines {
		if line[:3] == "you" {
			you = i
		}
		if line[:3] == "svr" {
			svr = i
		}
		if line[:3] == "fft" {
			fft = i
		}
		if line[:3] == "dac" {
			dac = i
		}
		names[i] = line[:3]
	}
	for i, line := range lines {
		for _, conn := range strings.Split(line[5:], " ") {
			connections[i] = append(connections[i], slices.Index(names, conn))
		}

	}
	//fmt.Println(you, out, svr, fft, dac)
	//fmt.Println(connections)
	//fmt.Println(names)

	srv2dac := countExit(svr, dac, nil)
	dac2fft := countExit(dac, fft, nil)
	fft2out := countExit(fft, out, nil)
	srv2fft := countExit(svr, fft, nil)
	fft2dac := countExit(fft, dac, nil)
	dac2out := countExit(dac, out, nil)

	fmt.Println(srv2dac*dac2fft*fft2out + srv2fft*fft2dac*dac2out)
}
