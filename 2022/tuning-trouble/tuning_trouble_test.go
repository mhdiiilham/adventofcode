package tuningtrouble_test

import (
	"fmt"
	"testing"

	tuningtrouble "github.com/mhdiiilham/adventofcode/2022/tuning-trouble"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const input = "../../input/2022/tuning_trouble_test.txt"

type tuningTroubleTestSuite struct {
	suite.Suite
}

func TestTuningTrouble(t *testing.T) {
	suite.Run(t, new(tuningTroubleTestSuite))
}

func (suite *tuningTroubleTestSuite) TestLoadInput() {
	testCases := []struct {
		name       string
		input      string
		dataStream string
		expeceted  error
	}{
		{
			name:       "success",
			input:      input,
			dataStream: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expeceted:  nil,
		},
		{
			name:       "failed",
			input:      "hello.txt",
			dataStream: "",
			expeceted:  fmt.Errorf("failed to read input file %s", "hello.txt"),
		},
	}

	t := suite.T()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tuningtrouble := tuningtrouble.New(tt.input)
			actual := tuningtrouble.LoadInput()
			assert.Equal(t, tt.expeceted, actual)
			assert.Equal(t, tt.dataStream, tuningtrouble.Datastream)
		})
	}
}

func (suite *tuningTroubleTestSuite) TestGetHowManyCharactersNeedToBeProcessedBeforeTheFirstStartOfPacketMarkerDetected() {
	testCase := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "success", input: "../../input/2022/tuning_trouble_test_0.txt", expected: 0},
		{name: "success", input: "../../input/2022/tuning_trouble_test.txt", expected: 7},
		{name: "success", input: "../../input/2022/tuning_trouble_test_2.txt", expected: 5},
		{name: "success", input: "../../input/2022/tuning_trouble_test_3.txt", expected: 10},
		{name: "success", input: "../../input/2022/tuning_trouble_test_4.txt", expected: 11},
	}

	t := suite.T()
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			assertion := assert.New(t)

			tuningTrouble := tuningtrouble.New(tt.input)
			err := tuningTrouble.LoadInput()
			assertion.NoError(err)

			actual := tuningTrouble.GetHowManyCharactersNeedToBeProcessedBeforeTheFirstStartOfPacketMarkerDetected()
			assertion.Equal(tt.expected, actual)
		})
	}
}

func (suite *tuningTroubleTestSuite) TestGetHowManyCharactersNeedToBeProcessedBeforeTheFirstStartOfMessageDetected() {
	testCase := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "success", input: "../../input/2022/tuning_trouble_test_0.txt", expected: 0},
		{name: "success", input: "../../input/2022/tuning_trouble_test.txt", expected: 19},
	}

	t := suite.T()
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			assertion := assert.New(t)

			tuningTrouble := tuningtrouble.New(tt.input)
			err := tuningTrouble.LoadInput()
			assertion.NoError(err)

			actual := tuningTrouble.GetHowManyCharactersNeedToBeProcessedBeforeTheFirstStartOfMessageDetected()
			assertion.Equal(tt.expected, actual)
		})
	}
}
