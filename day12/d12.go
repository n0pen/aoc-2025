package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("day12/input.txt")
	reg := regexp.MustCompile(`^\d+:$`)
	lines := strings.Split(string(file), "\n")
	var shapes []uint
	i := 0
	for ; ; i++ {
		if reg.MatchString(lines[i]) {
			i += 1
			val := uint(0)
			for r := 0; r < 3; r++ {
				for _, ch := range lines[i] {
					if ch == '#' {
						//val |= 1 << (j + r*3)
						val += 1
					}
				}
				i += 1
			}
			shapes = append(shapes, val)
		} else {
			break
		}

	}
	fmt.Printf("%b\n", shapes)
	sum := 0
	for ; i < len(lines); i++ {
		line := lines[i]
		h, _ := strconv.ParseInt(line[0:2], 10, 64)
		w, _ := strconv.ParseInt(line[3:5], 10, 64)
		vals := strings.Split(line[7:], " ")
		fmt.Println(vals)
		area := uint(0)
		for j := 0; j < len(shapes); j++ {
			ammount, _ := strconv.ParseInt(vals[j], 10, 64)
			area += uint(ammount) * shapes[j]

		}
		if uint(h*w) >= area {
			sum += 1
		}
	}
	fmt.Println(sum)

}
