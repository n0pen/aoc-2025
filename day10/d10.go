package main

import (
	"fmt"
	"math"
	"math/bits"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Diagram struct {
	lights  uint
	buttons []uint
}

func (diag Diagram) String() string {
	return fmt.Sprintf("lights: %010b, buttons: %010b", diag.lights, diag.buttons)
}

func toPresses(val uint) (count int) {
	for i := 0; i < bits.UintSize; i++ {
		count += int(val & 1)
		val >>= 1
	}
	return
}

func findMinButtonPresses(diag Diagram) (mini int) {
	mini = math.MaxInt
	search := uint(1<<(len(diag.buttons)) - 1)
	for i := uint(0); i <= search; i++ {
		if mini <= toPresses(i) {
			continue
		}
		lights := uint(0)
		for j := 0; j < len(diag.buttons); j++ {
			if (1<<j)&i == 1<<j {
				lights ^= diag.buttons[j]
			}
		}
		if lights == diag.lights {
			fmt.Printf("result %010b n:%d\n", i, toPresses(i))
			mini = min(mini, toPresses(i))
		}
	}
	return
}

func main() {
	redFile, _ := os.ReadFile("day10/input.txt")
	file := strings.Split(string(redFile), "\n")
	reg := regexp.MustCompile(`\[(?P<lights>.+)] (?P<buttons>(\(\d+(,\d+)*\)\s?)+) \{(?P<joltages>\d+(,\d+)*)}`)
	lightsIndex := reg.SubexpIndex("lights")
	buttonsIndex := reg.SubexpIndex("buttons")
	diagrams := make([]Diagram, len(file))
	sum := 0
	for a, line := range file {
		result := reg.FindStringSubmatch(line)
		lightsS, buttonsS := result[lightsIndex], result[buttonsIndex]
		lightsVal := uint(0)
		operations := make([]uint, strings.Count(buttonsS, " ")+1)

		for i, char := range lightsS {
			if char == '#' {
				lightsVal |= 1 << i
			}
		}
		for i, buts := range strings.Split(buttonsS, " ") {
			for _, val := range strings.Split(buts[1:len(buts)-1], ",") {
				parseInt, _ := strconv.ParseInt(val, 10, 64)
				operations[i] |= 1 << parseInt
			}
		}
		diagrams[a] = Diagram{lightsVal, operations}

		mini := findMinButtonPresses(diagrams[a])
		fmt.Printf("%+v %d\n", diagrams[a], mini)
		sum += mini

	}
	fmt.Printf("sum: %d\n", sum)

}
