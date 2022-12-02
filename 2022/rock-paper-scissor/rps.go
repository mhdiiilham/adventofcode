package rockpaperscissor

import (
	"strings"

	"github.com/mhdiiilham/adventofcode/pkg/input"
)

// X -> rock
// Y -> paper
// Z -> scissors
var (
	Strenght = map[string]map[string]bool{
		"X": {
			"Y": false,
			"Z": true,
		},
		"Y": {
			"X": true,
			"Z": false,
		},
		"Z": {
			"Z": false,
			"Y": true,
		},
	}

	ShapeScore = map[string]int{"X": 1, "Y": 2, "Z": 3}
)

type RockPaperScissor struct {
	InputSource string
	Input       [][]string
	Rows        []string
	Score       int
}

func NewRockPaperScissors(inputSource string) *RockPaperScissor {
	return &RockPaperScissor{InputSource: inputSource}
}

func (rps *RockPaperScissor) GetTotalScore() (score int) {
	for _, round := range rps.Input {
		enemy := round[0]
		myself := round[1]

		shapeScore := ShapeScore[myself]
		strenghtScore := 0

		if myself == enemy {
			strenghtScore = 3
		} else if Strenght[myself][enemy] {
			strenghtScore = 6
		}

		totalScore := shapeScore + strenghtScore
		rps.Score += totalScore
	}

	return rps.Score
}

func (rps *RockPaperScissor) LoadInput() (*RockPaperScissor, error) {
	var err error
	rps.Rows, err = input.ReadByLines(rps.InputSource, "\n")
	if err != nil {
		return rps, err
	}

	for _, row := range rps.Rows {
		enemyDict := map[string]string{"A": "X", "B": "Y", "C": "Z"}
		inputChallenge := strings.Split(row, " ")
		enemy := inputChallenge[0]

		strategy := []string{enemyDict[enemy], inputChallenge[1]}
		rps.Input = append(rps.Input, strategy)
	}

	return rps, nil
}

func (rps *RockPaperScissor) CalculateRealInput() int {
	totalScore := 0
	dict := map[string]map[string]int{
		"X": {
			"A": 3,
			"B": 1,
			"C": 2,
		}, // LOSE
		"Y": {
			"A": 1 + 3,
			"B": 2 + 3,
			"C": 3 + 3,
		}, // DRAW
		"Z": {
			"A": 2 + 6,
			"B": 3 + 6,
			"C": 1 + 6,
		}, // WIN
	}

	for _, row := range rps.Rows {
		value := strings.Split(row, " ")
		enemy := value[0]
		action := value[1]

		score := dict[action][enemy]
		totalScore += score
	}

	return totalScore
}
