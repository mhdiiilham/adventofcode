package campcleanup_test

import (
	"testing"

	campcleanup "github.com/mhdiiilham/adventofcode/2022/camp-cleanup"
	"github.com/stretchr/testify/assert"
)

const input = "../../input/2022/camp_cleanup_test.txt"

func TestGetNumbersOfAssignmentThatFullyContainsTheOtherPairAssignment(t *testing.T) {
	assertion := assert.New(t)
	expected := 2

	cc := campcleanup.NewCampCleanUp(input)
	errLoadInput := cc.LoadInput()
	assertion.NoError(errLoadInput)

	actual := cc.GetNumbersOfAssignmentThatFullyContainsTheOtherPairAssignment()
	assertion.Equal(expected, actual)
}

func TestGetNumbersOfAssignmentPairsDoTheRangeOverlap(t *testing.T) {
	assertion := assert.New(t)
	expected := 4

	cc := campcleanup.NewCampCleanUp(input)
	errLoadInput := cc.LoadInput()
	assertion.NoError(errLoadInput)

	actual := cc.GetNumbersOfAssignmentPairsDoTheRangeOverlap()
	assertion.Equal(expected, actual)
}

func TestInputNotFound(t *testing.T) {
	cc := campcleanup.NewCampCleanUp("")
	errLoadInput := cc.LoadInput()
	assert.Error(t, errLoadInput)
}
