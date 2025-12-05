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
	sum := 0

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
	ids := strings.Split(input[1], "\n")
	ids2 := make([]int64, len(ids))

	for i, line := range ids {
		val, _ := strconv.ParseInt(line, 10, 64)
		ids2[i] = val
	}
	slices.SortFunc(ids2, func(i, j int64) int { return int(i - j) })

	lastRage := 0
	fmt.Println(ids2)
	fmt.Println(ranges2)
	for i := 0; i < len(ids2); i++ {

		for j := lastRage; j < len(ranges2); j++ {
			if ids2[i] >= ranges2[j].min {
				lastRage = j

				if ids2[i] <= ranges2[j].max {
					sum += 1
					break
				}
				if j+1 < len(ranges2) && ranges2[j+1].min > ids2[i] {

					break
				}
			}
		}

	}

	fmt.Println(sum)

}
