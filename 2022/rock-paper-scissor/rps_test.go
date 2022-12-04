package rockpaperscissor_test

import (
	"testing"

	rockpaperscissor "github.com/mhdiiilham/adventofcode/2022/rock-paper-scissor"
	"github.com/stretchr/testify/assert"
)

const inputSource = "../../input/2022/rock_paper_scissors_test.txt"

func TestGetTotalScore(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assertion := assert.New(t)
		expected := 15

		rps, err := rockpaperscissor.NewRockPaperScissors(inputSource).LoadInput()
		assertion.NoError(err)

		totalScore := rps.GetTotalScore()
		assertion.Equal(expected, totalScore)
	})
}

func TestFixInputGetTotalScore(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assertion := assert.New(t)
		expected := 12

		rps, err := rockpaperscissor.NewRockPaperScissors(inputSource).LoadInput()
		assertion.NoError(err)

		totalScore := rps.CalculateRealInput()
		assertion.Equal(expected, totalScore)
	})
}

func TestInputNotFound(t *testing.T) {
	_, errLoadInput := rockpaperscissor.NewRockPaperScissors("").LoadInput()
	assert.Error(t, errLoadInput)
}
