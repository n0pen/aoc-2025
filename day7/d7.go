package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func moveTachion(input []string) int {
	sum := 0
	nextString := []byte(input[1])
	start := slices.Index([]byte(input[0]), 'S')
	nextString[start] = '|'
	regBeam := regexp.MustCompile("\\|")
	for i := 2; i < len(input)-1; i++ {
		beams := regBeam.FindAllIndex(nextString, -1)
		nextString = []byte(input[i])
		for _, arr := range beams {
			if input[i][arr[0]] == '^' {
				sum += 1
				nextString[arr[0]-1] = '|'
				nextString[arr[0]+1] = '|'
			} else {
				nextString[arr[0]] = '|'
			}
		}
	}
	return sum
}

func main() {

	file, _ := os.ReadFile("day7/input.txt")
	input := strings.Split(string(file), "\n")

	fmt.Println(moveTachion(input))

}
