package main

import (
	"fmt"
	"os"
	"strings"
)

func get_space(x, y int, storage []string) int {
	masks := []string{
		"@@@", "@ @", "@@@",
	}
	count := 8
	for posy := 0; posy < len(masks); posy++ {

		scany := posy - 1 + y
		if scany < 0 || scany >= len(storage) {
			continue
		}
		for posx := 0; posx < len(masks[posy]); posx++ {
			scanx := posx - 1 + x
			if scanx < 0 || scanx >= len(storage[y]) {
				continue
			}
			//fmt.Println(storage[scany][scanx], masks[posy][posx])
			if storage[scany][scanx] == masks[posy][posx] {
				count -= 1
			}
		}
	}
	if count > 4 {
		return 1
	}
	return 0
}

func main() {
	file, _ := os.ReadFile("./day4/input.txt")
	sum := 0

	str := strings.Split(string(file), "\n")

	fmt.Println(get_space(5, 2, str))

	for i, row := range str {
		for j, _ := range row {
			if row[j] == '@' {
				sum += get_space(j, i, str)
			}
		}
	}
	fmt.Println(len(str), sum)
}
