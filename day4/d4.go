package main

import (
	"fmt"
	"os"
	"strings"
)

func getSpace(x, y int, storage []string) int {
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

	for {
		lpcnt := 0
		for i, row := range str {
			for j := range row {
				if row[j] == '@' {
					res := getSpace(j, i, str)
					sum += res
					lpcnt += res
					if res > 0 {
						var val = &str[i]
						*val = str[i][0:j] + " " + str[i][j+1:]
					}
				}
			}
		}
		if lpcnt == 0 {
			break
		}
	}
	fmt.Println(sum)
}
