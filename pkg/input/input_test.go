package input_test

import (
	"testing"

	"github.com/mhdiiilham/adventofcode/pkg/input"
	"github.com/stretchr/testify/assert"
)

func TestReadByLines(t *testing.T) {

	t.Run("success: seperator one line", func(t *testing.T) {
		inputFile := "sample_1.txt"
		expectedResult := []string{"1", "2", "3"}

		assertion := assert.New(t)
		result, err := input.ReadByLines(inputFile, input.OneEnter)
		assertion.Equal(expectedResult, result)
		assertion.NoError(err)
	})

	t.Run("failed: file not found", func(t *testing.T) {
		inputFile := "sample_1_not_found.txt"

		assertion := assert.New(t)
		result, err := input.ReadByLines(inputFile, input.OneEnter)
		assertion.Empty(result)
		assertion.Error(err)
	})
}
