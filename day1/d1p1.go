package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const dialLength = 100

func main() {
	position := 50
	file, _ := os.ReadFile("./day1/input")
	counter := 0
	filestr := strings.Replace(string(file), "\r\n", "\n", -1)
	for _, str := range strings.Split(filestr, "\n") {

		str = strings.Replace(str, "L", "-", -1)
		str = strings.Replace(str, "R", "", -1)
		rotation, _ := strconv.ParseInt(str, 10, 64)
		times := int(math.Abs(float64(int(rotation) / dialLength)))
		counter += times
		newval := int(rotation) % dialLength
		newPosition := position + newval

		if newPosition > dialLength || (position > 0 && newPosition < 0) || newPosition%dialLength == 0 {
			counter += 1
		}
		position = int(math.Abs(float64((newPosition + dialLength) % dialLength)))

	}
	fmt.Println(counter)

}
