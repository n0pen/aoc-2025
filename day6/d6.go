package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./day6/input.txt")
	sum := int64(0)
	var vals [][]int64

	input := strings.Split(string(file), "\n")
	if len(input) < 2 {
		fmt.Println(0)
		return
	}
	vals = make([][]int64, len(input)-1)
	var operators []string

	for i, row := range input {
		reg := regexp.MustCompile("\\s+")
		rowvals := reg.Split(row, -1)

		if i < len(vals) {
			vals[i] = make([]int64, len(rowvals))
		} else {
			operators = make([]string, len(rowvals))
		}
		for j, valstring := range rowvals {
			val := strings.ReplaceAll(valstring, " ", "")
			if i < len(vals) {

				vals[i][j], _ = strconv.ParseInt(val, 10, 64)
			} else {
				operators[j] = val
			}
		}
	}

	for i := 0; i < len(operators); i++ {

		rowval := vals[0][i]
		for j := 1; j < len(vals); j++ {
			if operators[i] == "*" {
				rowval *= vals[j][i]
			} else {
				rowval += vals[j][i]
			}
		}
		sum += rowval
		fmt.Println(sum)
	}

	fmt.Println(sum)
}
