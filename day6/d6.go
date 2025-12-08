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

	input := strings.Split(string(file), "\n")
	if len(input) < 2 {
		fmt.Println(0)
		return
	}
	var operators = input[len(input)-1]
	fmt.Println(operators)
	curind := 0
	for {

		reg := regexp.MustCompile("([*+])")
		op := operators[curind]
		lengarr := reg.FindIndex([]byte(operators[curind+1:]))
		var leng int
		if len(lengarr) == 2 {
			leng = lengarr[0]
		}
		var problem int64
		if op == '*' {
			problem = 1
		}

		fmt.Println(string(op), leng)
		var bonus int
		if leng == 0 {
			bonus = len(input[0]) - curind
		}
		for i := 0; i < leng+bonus; i++ {
			varstr := ""
			for j := 0; j < len(input)-1; j++ {
				varstr += string(input[j][curind+i])
			}
			fmt.Println(varstr)
			val, _ := strconv.ParseInt(strings.ReplaceAll(varstr, " ", ""), 10, 64)
			if op == '*' {
				problem *= val
			} else {
				problem += val
			}
		}
		sum += int64(problem)

		if leng == 0 {
			break
		}

		curind = curind + leng + 1

	}

	fmt.Println(sum)
}
