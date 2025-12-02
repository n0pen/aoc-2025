package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isProblem(value int64) bool {
	length := int64(math.Ceil(math.Log10(float64(value))))
	if length%2 == 1 {
		return false
	}
	pow := int64(math.Pow10(int(length >> 1)))
	return value/pow == value%pow
}

func main() {
	file, _ := os.ReadFile("./day2/input.txt")
	sum := int64(0)
	for _, str := range strings.Split(string(file), ",") {
		fmt.Println(str)
		res := strings.Split(str, "-")
		va0, err1 := strconv.ParseInt(res[0], 10, 64)
		if err1 != nil {
			panic(err1)
		}
		va1, _ := strconv.ParseInt(res[1], 10, 64)

		for counter := va0; counter <= va1; {
			if isProblem(counter) {
				sum += counter
			}
			counter++
		}

	}
	fmt.Println(sum)
}
