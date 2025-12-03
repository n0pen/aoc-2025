package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const length = 12

func main() {
	file, _ := os.ReadFile("./day3/input.txt")
	sum := int64(0)

	for _, str := range strings.Split(string(file), "\n") {
		values := make([]byte, length)
		lastIndex := 0
		fmt.Println(str)
		for i := length; i > 0; i-- {
			slice := strings.Split(str[lastIndex:len(str)-i+1], "")
			sort.Strings(slice)
			fmt.Println(str[lastIndex:len(str)-i], " ", slice[len(slice)-1])
			fmt.Println(strings.Join(slice, ""))
			value := slice[len(slice)-1]
			lastIndex = strings.Index(str[lastIndex:], value) + lastIndex + 1
			values[length-i] = value[0]
		}

		val, _ := strconv.ParseInt(string(values), 10, 64)
		fmt.Println(val)
		sum += val
	}
	fmt.Println(sum)
}
