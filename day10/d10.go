package main

import (
	"fmt"
	"math"
	"math/bits"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Diagram struct {
	lights  uint
	buttons []uint
	joltage []uint
}

func (diag Diagram) String() string {
	return fmt.Sprintf("lights: %010b, buttons: %010b, joltage: %d", diag.lights, diag.buttons, diag.joltage)
}

func toPresses(val uint) (count int) {
	for i := 0; i < bits.UintSize; i++ {
		count += int(val & 1)
		val >>= 1
	}
	return
}

func printMatrix[T float64](matrix [][]T, format string) (res string) {

	for i, val := range matrix {
		for j, t := range val {
			if j == len(matrix[i])-1 {
				if i < len(matrix)-1 && matrix[i][j] > 0 {
					res += fmt.Sprint("\t*")
				}
			} else {
				if j == len(matrix[i])-2 {
					res += "║"
				}
				res += fmt.Sprintf(format, t)

			}

		}
		if i == len(matrix)-2 {
			res += "\n"
			for j := 0; j < len(matrix[i]); j++ {

				res += "══════"
			}
		}
		res += "\n"
	}
	return
}

/*
returns the smallest non-zero joltage index
*/
func pressButton(button, times uint, vals []uint) uint {
	for i := 0; i < len(vals); i++ {
		if button&(1<<i) == 1<<i {
			vals[i] += times
		}
	}
	return 0
}

func (diag Diagram) createSimplex() [][]float64 {
	matrix := make([][]float64, len(diag.joltage)*2+1)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]float64, len(diag.buttons)+len(diag.joltage)*2+2)
	}
	for i, jolt := range diag.joltage {
		for j, button := range diag.buttons {
			if button&(1<<i) == 1<<i {
				matrix[i][j] = 1
				matrix[i+len(diag.joltage)][j] = 1
			}
			if i == 0 {
				matrix[len(diag.joltage)*2][j] = 1
			}
		}
		matrix[i][i+len(diag.buttons)] = -1
		matrix[i+len(diag.joltage)][i+len(diag.buttons)+len(diag.joltage)] = 1
		matrix[i][len(diag.buttons)+len(diag.joltage)*2] = float64(jolt)
		matrix[i+len(diag.joltage)][len(diag.buttons)+len(diag.joltage)*2] = float64(jolt)
		matrix[i][len(diag.buttons)+len(diag.joltage)*2+1] = 1

	}

	return matrix
}

func hasSolution(matrix [][]float64) bool {
	cols := len(matrix[0])
	rows := len(matrix)

	for j := 0; j < cols-2; j++ {
		if matrix[rows-1][j] < -0.00001 {
			return false
		}
	}

	return true
}

func canContinue(matrix [][]float64) bool {
	columns := len(matrix[0])
	rows := len(matrix)
	star := false
	for i := 0; i < rows-1; i++ {
		if matrix[i][columns-1] > 0 && matrix[i][columns-2] != 0 {
			star = true
			for j := 0; j < columns-2; j++ {
				if matrix[i][j] > 0.0001 {
					return true
				}
			}
		}
		if star {
			return false
		}
	}

	for i := 0; i < columns-2; i++ {
		if matrix[rows-1][i] < -0.0001 {
			return true
		}
	}
	return false
}

func iterateSimplex(matrix [][]float64) {
	columns := len(matrix[0])
	rows := len(matrix)
	//clean zeros
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] < 0.000001 && matrix[i][j] > -0.000001 {
				matrix[i][j] = 0
			}
		}
	}
	//remove extra stars
	for i := 0; i < rows; i++ {
		if matrix[i][columns-1] > 0 && matrix[i][columns-2] == 0 {
			matrix[i][columns-2] = 0
			matrix[i][columns-1] = 0
			for j := 0; j < columns-2; j++ {
				if matrix[i][j] != 0 {
					matrix[i][j] *= -1
				}

			}
		}
	}
	// find pivot
	pivc := -1
	for i := 0; i < rows-1; i++ {
		if matrix[i][columns-1] > 0 {
			//fmt.Printf("found star: %v ", i)
			for j := 0; j < columns-2; j++ {
				if matrix[i][j] > 0.00001 {
					if pivc < 0 {
						pivc = j
					} else {
						if matrix[i][pivc] < matrix[i][j] {
							pivc = j
						}
					}
				}

			}
			if pivc >= 0 {
				break
			}
		}
	}
	// no stars left
	if pivc == -1 {

		for j := 0; j < columns-2; j++ {
			if matrix[rows-1][j] < 0 {
				if pivc == -1 {
					pivc = j
				} else if matrix[rows-1][j] < matrix[rows-1][pivc] {
					pivc = j
				}
			}

		}
		if pivc == -1 {
			return
		}
	}
	pivr := -1
	for j := rows - 2; j >= 0; j-- {
		if matrix[j][pivc] > 0 && pivr < 0 {
			pivr = j
		} else if matrix[j][pivc] > 0 {
			if matrix[pivr][columns-2]/matrix[pivr][pivc] == matrix[j][columns-2]/matrix[j][pivc] {
				if matrix[pivr][columns-1] == 0 && matrix[j][columns-1] == 1 {
					pivr = j
				}
			} else if matrix[pivr][columns-2]/matrix[pivr][pivc] > matrix[j][columns-2]/matrix[j][pivc] {
				pivr = j
			}
		}
	}
	if pivr == -1 {

		return
	}

	//fmt.Printf("piv: r %v, c %v : %v ratio %v\n", pivr, pivc, matrix[pivr][pivc], matrix[pivr][columns-2])
	if matrix[pivr][pivc] != 1 {
		val := matrix[pivr][pivc]
		for i := 0; i < columns-1; i++ {
			matrix[pivr][i] /= val
		}
	}
	red1 := matrix[pivr][pivc]

	for j := 0; j < len(matrix); j++ {
		if j == pivr || matrix[j][pivc] == 0 {
			continue
		}
		blue1 := matrix[j][pivc]
		for w := 0; w < len(matrix[j])-1; w++ {
			r2 := matrix[j][w]
			r1 := matrix[pivr][w]
			mult := float64(1)
			if red1*blue1 > 0 {
				mult = -1
			}

			matrix[j][w] = r2*max(red1, -red1) + mult*r1*max(blue1, -blue1)

		}
	}
	matrix[pivr][columns-1] = 0

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			val := matrix[i][j]
			if val > 0 {
				if math.Ceil(val)-val < 0.00001 {
					matrix[i][j] = math.Ceil(val)
				} else if val-math.Floor(val) < 0.00001 {
					matrix[i][j] = math.Floor(val)
				}
			}

		}
	}

}

func printAsProblem(matrix [][]float64) {
	fmt.Print("Minimize p=")
	columns := len(matrix[0])
	rows := len(matrix)
	fmt.Print("xA")
	for i := 1; i < columns-rows-1; i++ {
		fmt.Printf(" + x%v", string([]byte{byte('A' + i)}))
	}
	fmt.Println()
	for i := 0; i < rows-1; i++ {
		fmt.Printf("%vxA", int(matrix[i][0]))
		for j := 1; j < columns-rows-1; j++ {
			fmt.Printf(" + %vx%v", int(matrix[i][j]), string([]byte{byte('A' + j)}))
		}
		sign := "<"
		if matrix[i][columns-1] > 0 {
			sign = ">"
		}
		fmt.Printf("%v=%v\n", sign, int(matrix[i][columns-2]))
	}
}

func getTotal(matrix [][]float64, diag Diagram) (sum int) {
	if !hasSolution(matrix) {
		return math.MaxInt
	}
	buttons := len(diag.buttons)
	presses := getPresses(matrix, diag)
	//fmt.Printf("press: %v\n", presses)
	//presses = []float64{1, 5, 0, 1, 3, 0}
	mini := uint(math.MaxInt)

	values := make([]uint, len(diag.joltage))
	pressesI := make([]uint, len(diag.buttons))
	for j := 0; j < buttons; j++ {
		if presses[j]-math.Floor(presses[j]) > 0.0001 || math.Ceil(presses[j])-presses[j] > 0.0001 {
			return math.MaxInt
		}
		times := uint(presses[j])
		pressesI[j] = times
		pressButton(diag.buttons[j], times, values)
	}
	//fmt.Printf("vals: %v jolts: %v press: %v\n", values, diag.joltage, pressesI)
	if slices.Compare(values, diag.joltage) == 0 {

		if s := sumAll(pressesI); s < mini {
			mini = s
		}
	}

	sum = int(mini)
	//fmt.Println(sum)
	return
}

func getPresses(matrix [][]float64, diag Diagram) []float64 {
	presses := make([]float64, len(diag.buttons))
	cols := len(matrix[0])
	rows := len(matrix)

	for i := 0; i < len(presses); i++ {
		candidate := -1
		for j := 0; j < rows-1; j++ {
			if matrix[j][i] != 0 {
				if candidate == -1 {
					if matrix[j][i] == 1 {

						candidate = j
					} else {
						break
					}
				} else {
					candidate = -1
					break
				}
			}
		}
		if candidate != -1 {
			presses[i] = matrix[candidate][cols-2]
		}
	}
	return presses
}
func sumAll(arr []uint) (s uint) {

	for _, val := range arr {
		s += val

	}
	return
}

var branching = 10

type Constraint struct {
	idx, value int
}

func branchSimplex(diagram Diagram, optimal, initial [][]float64, constraints *[]Constraint) int {

	branching -= 1
	if branching < 0 {
		//return -1
	}
	/*fmt.Println("Branching")
	printAsProblem(initial)
	fmt.Println("optimal")
	fmt.Println(printMatrix(optimal, "%6.1f"))
	fmt.Println("initial")
	fmt.Println(printMatrix(initial, "%6.1f"))*/
	presses := getPresses(optimal, diagram)
	var pivots []int
	maxdifidx := -1
	for i, press := range presses {
		if diff := press - math.Floor(press); diff > 0.000001 {
			con := Constraint{i, int(math.Floor(press))}
			if slices.Index(*constraints, con) >= 0 {
				continue
			}
			*constraints = append(*constraints, con)
			pivots = append(pivots, i)
			if maxdifidx == -1 {
				maxdifidx = i
			} else if diff > presses[maxdifidx]-math.Floor(presses[maxdifidx]) {
				maxdifidx = i
			}
		}
	}
	if maxdifidx == -1 {
		return getTotal(optimal, diagram)
	}
	mini := math.MaxInt
	for _, piv := range pivots {

		rows := len(initial)
		cols := len(initial[0])
		s1 := make([][]float64, rows)
		s2 := make([][]float64, rows)
		s11 := make([][]float64, rows)
		s22 := make([][]float64, rows)
		for i := 0; i < rows; i++ {
			s1[i] = make([]float64, cols)
			s11[i] = make([]float64, cols)
			s2[i] = make([]float64, cols)
			s22[i] = make([]float64, cols)
			copy(s1[i], initial[i])
			copy(s2[i], initial[i])
			copy(s11[i], initial[i])
			copy(s22[i], initial[i])

			s1[i] = append(s1[i][:cols-2], 0, s1[i][cols-2], s1[i][cols-1])
			s2[i] = append(s2[i][:cols-2], 0, s2[i][cols-2], s2[i][cols-1])

			s11[i] = append(s11[i][:cols-2], 0, s11[i][cols-2], s11[i][cols-1])
			s22[i] = append(s22[i][:cols-2], 0, s22[i][cols-2], s22[i][cols-1])
		}
		s1 = append(s1[:rows-1], make([]float64, cols+1), s1[rows-1])
		s2 = append(s2[:rows-1], make([]float64, cols+1), s2[rows-1])
		s1[rows-1][piv] = 1
		s1[rows-1][cols-2] = -1
		s1[rows-1][cols-1] = math.Ceil(presses[piv])
		s1[rows-1][cols] = 1
		s2[rows-1][piv] = 1
		s2[rows-1][cols-2] = 1
		s2[rows-1][cols-1] = math.Floor(presses[piv])
		s11 = append(s11[:rows-1], make([]float64, cols+1), s11[rows-1])
		s22 = append(s22[:rows-1], make([]float64, cols+1), s22[rows-1])
		s11[rows-1][piv] = 1
		s11[rows-1][cols-2] = -1
		s11[rows-1][cols-1] = math.Ceil(presses[piv])
		s11[rows-1][cols] = 1
		s22[rows-1][piv] = 1
		s22[rows-1][cols-2] = 1
		s22[rows-1][cols-1] = math.Floor(presses[piv])

		//fmt.Println("Branching 1")
		//fmt.Println(printMatrix(s1, "%6.1f"))
		solveSimplex(s1)
		//fmt.Println("Branching 2")
		//fmt.Println(printMatrix(s2, "%6.1f"))
		solveSimplex(s2)
		//fmt.Printf("totals: s1 %v, s2 %v\n", s1[rows][cols-1], s2[rows][cols-1])
		val1, val2 := math.MaxInt, math.MaxInt
		if !hasSolution(s1) && !hasSolution(s2) {
			continue
		}
		if hasSolution(s2) {
			//fmt.Println("digging into s2")
			val1 = branchSimplex(diagram, s2, s22, constraints)
		}

		if val1 < mini {
			mini = val1
		}
		if hasSolution(s1) {
			//fmt.Println("digging into s1")
			val2 = branchSimplex(diagram, s1, s11, constraints)
		}

		if val2 < mini {
			mini = val2
		}
	}
	return mini
}

func main() {
	redFile, _ := os.ReadFile("day10/input.txt")
	file := strings.Split(string(redFile), "\n")
	reg := regexp.MustCompile(`\[(?P<lights>.+)] (?P<buttons>(\(\d+(,\d+)*\)\s?)+) \{(?P<joltages>\d+(,\d+)*)}`)
	lightsIndex := reg.SubexpIndex("lights")
	buttonsIndex := reg.SubexpIndex("buttons")
	joltagesIndex := reg.SubexpIndex("joltages")
	diagrams := make([]Diagram, len(file))
	sum := 0
	for a, line := range file {
		result := reg.FindStringSubmatch(line)
		lightsS, buttonsS, joltagesS := result[lightsIndex], result[buttonsIndex], result[joltagesIndex]
		lightsVal := uint(0)
		operations := make([]uint, strings.Count(buttonsS, " ")+1)
		joltages := make([]uint, len(lightsS))

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

		for i, jolt := range strings.Split(joltagesS, ",") {
			val, _ := strconv.ParseInt(jolt, 10, 64)
			joltages[i] = uint(val)
		}

		diagrams[a] = Diagram{lightsVal, operations, joltages}
		fmt.Println(diagrams[a])
		simplex := diagrams[a].createSimplex()

		/*printAsProblem(simplex)
		fmt.Println(printMatrix(simplex, "%6.1f"))*/
		solveSimplex(simplex)
		//fmt.Println(printMatrix(simplex, "%6.1f"))

		total := getTotal(simplex, diagrams[a])
		if total == -1 || total == math.MaxInt {
			//fmt.Printf("%v %v\n", a, total)
			total = branchSimplex(diagrams[a], simplex, diagrams[a].createSimplex(), &[]Constraint{})
			if total == -1 || total == math.MaxInt {
				panic("didn't find any solution")
			}
		}
		fmt.Println(total)
		//fmt.Println(a, res, total, res-float64(total))
		sum += total

	}
	fmt.Printf("sum: %d\n", sum)

}

func solveSimplex(simplex [][]float64) bool {
	for i := 0; canContinue(simplex); i++ {
		//fmt.Println(printMatrix(simplex, "%6.1f"))

		if i > 20 {
			//fmt.Println(printMatrix(simplex, "%6.1f"))
			//return false
			//panic("test")
		}
		if branching > 3 && i > 10 {
			//fmt.Println("iteration", i)
			//fmt.Println(printMatrix(simplex, "%6.1f"))
		}
		iterateSimplex(simplex)
		//fmt.Println(printMatrix(simplex, "%6.1f"))

	}
	if !hasSolution(simplex) {
		//fmt.Println("No solution found")
		//fmt.Println(printMatrix(simplex, "%6.1f"))
		return false
	}
	return true
}
