package rednosereport

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Solver hold the data needed to solve the Red-Nose Reports challenge.
type Solver struct {
	inputFilePath string
	reports       [][]int
}

// NewSolver return the Solver struct of challenge Red-Nose Reports.
// inputFilePath is the location of the puzzle input file.
func NewSolver(inputFilePath string) (*Solver, error) {
	s := &Solver{inputFilePath: inputFilePath}
	return s.read()
}

// Solve function solve the Red-Nose Reports challenge.
func (s *Solver) Solve() (int, int) {
	return s.partOne(), s.partTwo()
}

func (s *Solver) partOne() int {
	var numberOfSafeReport int

	for _, reports := range s.reports {
		reportsAreSafe := true
		increase := false
		decrease := false
		for i := 0; i < len(reports)-1; i++ {
			spike := reports[i] - reports[i+1]
			if spike > 0 {
				increase = true
			}

			if spike < 0 {
				decrease = true
			}

			if spike > 3 || spike < -3 || reports[i] == reports[i+1] || (increase && decrease) {
				reportsAreSafe = false
			}
		}

		if reportsAreSafe {
			numberOfSafeReport++
		}
	}

	return numberOfSafeReport
}

func (s *Solver) partTwo() int {
	var numberOfSafeReport int

	for _, reports := range s.reports {
		if s.checkIfReportsAreSave(reports) {
			numberOfSafeReport++
		} else {
			for i := 0; i < len(reports); i++ {
				if s.checkIfReportsAreSave(removeIndex(reports, i)) {
					numberOfSafeReport++
					break
				}
			}
		}
	}

	return numberOfSafeReport
}

func (s *Solver) checkIfReportsAreSave(reports []int) bool {
	reportsAreSafe := true
	increase := false
	decrease := false
	for i := 0; i < len(reports)-1; i++ {
		spike := reports[i] - reports[i+1]
		if spike > 0 {
			increase = true
		}

		if spike < 0 {
			decrease = true
		}

		if spike > 3 || spike < -3 || reports[i] == reports[i+1] || (increase && decrease) {
			reportsAreSafe = false
		}
	}

	return reportsAreSafe
}

func (s *Solver) read() (*Solver, error) {
	inputInBytes, err := os.ReadFile(s.inputFilePath)
	if err != nil {
		log.Fatalf("[Red-Nose Reports] failed to read input: %v", err)
		return s, err
	}

	reportsString := strings.Split(string(inputInBytes), "\n")
	for _, report := range reportsString {
		reportValue := []int{}
		for _, reportValueInStr := range strings.Split(report, " ") {
			v, _ := strconv.Atoi(reportValueInStr)
			reportValue = append(reportValue, v)
		}
		s.reports = append(s.reports, reportValue)

	}

	return s, nil
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
