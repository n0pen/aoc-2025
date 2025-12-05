package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type ValueRange struct {
	min int64
	max int64
}

func main() {
	file, _ := os.ReadFile("./day5/input.txt")
	sum := int64(0)

	input := strings.Split(string(file), "\n\n")

	ranges := strings.Split(input[0], "\n")
	ranges2 := make([]ValueRange, len(ranges))
	for i, line := range ranges {

		vals := strings.Split(line, "-")
		minv, _ := strconv.ParseInt(vals[0], 10, 64)
		maxv, _ := strconv.ParseInt(vals[1], 10, 64)
		val := ValueRange{
			minv,
			maxv,
		}

		ranges2[i] = val
	}
	slices.SortFunc(ranges2, func(i, j ValueRange) int { return int(i.min - j.min) })
	fmt.Println(ranges2)

	for j := 0; j < len(ranges2); j++ {

		curmax, curmin := ranges2[j].max, ranges2[j].min

		for i := j + 1; ; i++ {

			if i < len(ranges2) && ranges2[i].min <= curmax {
				if ranges2[i].max > curmax {
					curmax = ranges2[i].max
				}

				j += 1

			} else {
				sum += curmax - curmin + 1
				break
			}
		}

	}

	fmt.Println(sum)

}
