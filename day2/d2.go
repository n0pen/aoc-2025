package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isProblem(value int64) bool {

	str := strconv.FormatInt(value, 10)
	for i := 1; i <= len(str)/2; i++ {
		newstr := str[:i]
		cnt := strings.Count(str, newstr)
		if cnt*i >= len(str) {
			return true
		}
	}
	return false

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
