package campcleanup

import (
	"strconv"
	"strings"

	"github.com/mhdiiilham/adventofcode/pkg/input"
)

type CampCleanup struct {
	InputSource string
	Assignments [][]string
}

func NewCampCleanUp(inputSource string) *CampCleanup {
	return &CampCleanup{InputSource: inputSource}
}

func (c *CampCleanup) LoadInput() error {
	rows, err := input.ReadByLines(c.InputSource, input.OneEnter)
	if err != nil {
		return err
	}

	for _, row := range rows {
		c.Assignments = append(c.Assignments, strings.Split(row, ","))
	}

	return nil
}
func (c *CampCleanup) GetNumbersOfAssignmentThatFullyContainsTheOtherPairAssignment() int {
	var result int
	for _, assignment := range c.Assignments {
		assignmentOne := strings.Split(assignment[0], "-")
		assignmentTwo := strings.Split(assignment[1], "-")

g		assignmentOneRangeStart, assignmentOneRangeEnd := getRangeStartEnd(assignmentOne)
		assignmentTwoRangeStart, assignmentTwoRangeEnd := getRangeStartEnd(assignmentTwo)

		// With help from my favorite girl
		// https://www.tiktok.com/@richrachellll
		if assignmentOneRangeStart <= assignmentTwoRangeStart && assignmentOneRangeEnd >= assignmentTwoRangeEnd {
			result++
		} else if assignmentTwoRangeStart <= assignmentOneRangeStart && assignmentTwoRangeEnd >= assignmentOneRangeEnd {
			result++
		}
	}
	return result
}

func (c *CampCleanup) GetNumbersOfAssignmentPairsDoTheRangeOverlap() int {
	var result int
	for _, assignment := range c.Assignments {
		assignmentOne := strings.Split(assignment[0], "-")
		assignmentTwo := strings.Split(assignment[1], "-")

		assignmentOneRangeStart, assignmentOneRangeEnd := getRangeStartEnd(assignmentOne)
		assignmentTwoRangeStart, assignmentTwoRangeEnd := getRangeStartEnd(assignmentTwo)

		if assignmentOneRangeStart <= assignmentTwoRangeEnd && assignmentTwoRangeStart <= assignmentOneRangeEnd {
			result++
		}
	}
	return result
}

func getRangeStartEnd(assignment []string) (start, end int) {
	start, _ = strconv.Atoi(assignment[0])
	end, _ = strconv.Atoi(assignment[1])
	return
}
