package mullitover

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Solver struct hold required data to solve the Mull it Over challenge.
type Solver struct {
	inputFilePath         string
	curroptedInstructions []string
}

// NewSolver function return the Solver of Mull It Over challenge.
func NewSolver(inputFilePath string) (*Solver, error) {
	s := &Solver{inputFilePath: inputFilePath}
	return s.read()
}

// Solve function solve the Mull It Over challenge.
func (s *Solver) Solve() (int, int) {
	return s.solve(`mul\(\d+,\d+\)`), s.solve(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)
}

func (s *Solver) solve(expression string) int {
	var result int

	re := regexp.MustCompile(expression)
	mulEnabled := true

	for _, instruction := range s.curroptedInstructions {
		matches := re.FindAllStringSubmatch(instruction, -1)

		for _, match := range matches {
			if strings.HasPrefix(match[0], "don't") {
				mulEnabled = false
			} else if strings.HasPrefix(match[0], "do") {
				mulEnabled = true
			} else if strings.HasPrefix(match[0], "mul(") && mulEnabled {
				params := match[0][4 : len(match[0])-1]
				numbers := strings.Split(params, ",")
				if len(numbers) == 2 {
					x, _ := strconv.Atoi(numbers[0])
					y, _ := strconv.Atoi(numbers[1])
					result += x * y
				}
			}
		}
	}

	return result
}

func (s *Solver) read() (*Solver, error) {
	inputInBytes, err := os.ReadFile(s.inputFilePath)
	if err != nil {
		return s, err
	}

	curroptedInstructions := strings.Split(string(inputInBytes), "\n")
	s.curroptedInstructions = curroptedInstructions

	return s, nil
}
