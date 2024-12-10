package historianhysteria

import (
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Solver hold the data needed to solve the historian hysteria challenge.
type Solver struct {
	inputFilePath string
	rightList     []int
	leftList      []int
}

// NewSolver return the Solver struct of challenge Historian Hysteria.
// inputFilePath is the location of the puzzle input file.
func NewSolver(inputFilePath string) (*Solver, error) {
	s := &Solver{inputFilePath: inputFilePath}
	return s.read()
}

// Solve function solve the historian hysteria challenge.
func (s *Solver) Solve() (int, int) {
	return s.partOne(), s.partTwo()
}

func (s *Solver) partOne() int {
	var result int

	slices.Sort(s.leftList)
	slices.Sort(s.rightList)

	for idx := range s.leftList {
		Distance := math.Max(float64(s.leftList[idx]), float64(s.rightList[idx])) - math.Min(float64(s.leftList[idx]), float64(s.rightList[idx]))
		result += int(Distance)
	}

	return result
}

func (s *Solver) partTwo() int {
	var result int

	for _, id := range s.leftList {
		appearance := 0
		for _, v := range s.rightList {
			if id == v {
				appearance++
			}
		}

		result += id * appearance
	}

	return result
}

func (s *Solver) read() (*Solver, error) {
	inputInBytes, err := os.ReadFile(s.inputFilePath)
	if err != nil {
		return s, err
	}

	leftDistances := []int{}
	rightDistances := []int{}

	distances := strings.Split(string(inputInBytes), "\n")
	for _, distance := range distances {
		leftDistance, _ := strconv.Atoi(strings.Split(distance, "   ")[0])
		rightDistance, _ := strconv.Atoi(strings.Split(distance, "   ")[1])
		leftDistances = append(leftDistances, leftDistance)
		rightDistances = append(rightDistances, rightDistance)
	}

	s.leftList = leftDistances
	s.rightList = rightDistances

	return s, nil
}
