package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	position := 100000000 + 50
	file, _ := os.ReadFile("./day1/input")
	counter := 0
	filestr := strings.Replace(string(file), "\r\n", "\n", -1)
	for _, str := range strings.Split(filestr, "\n") {
		str = strings.Replace(str, "L", "-", -1)
		str = strings.Replace(str, "R", "", -1)
		val, _ := strconv.ParseInt(str, 10, 64)
		position += int(val)
		if position%100 == 0 {
			counter += 1
		}
	}
	fmt.Println(counter)

}
