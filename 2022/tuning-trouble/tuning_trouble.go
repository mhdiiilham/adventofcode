package tuningtrouble

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type TuningTrouble struct {
	inputSource string
	dataStream  string
}

const (
	packet  = 4
	message = 14
)

func New(inputSource string) *TuningTrouble {
	return &TuningTrouble{inputSource: inputSource}
}

func (tt *TuningTrouble) LoadInput() (err error) {
	b, err := ioutil.ReadFile(tt.inputSource)
	if err != nil {
		return fmt.Errorf("failed to read input file %s", tt.inputSource)
	}

	tt.dataStream = string(b)
	return nil
}

func (tt *TuningTrouble) GetHowManyCharactersNeedToBeProcessedBeforeTheFirstStartOfPacketMarkerDetected() int {
	return tt.getHowManyCharactersNeedToBeProcessedBeforeTheFirstStart(packet)
}

func (tt *TuningTrouble) GetHowManyCharactersNeedToBeProcessedBeforeTheFirstStartOfMessageDetected() int {
	return tt.getHowManyCharactersNeedToBeProcessedBeforeTheFirstStart(message)
}

func (tt *TuningTrouble) getHowManyCharactersNeedToBeProcessedBeforeTheFirstStart(numberOfDistinctChars int) int {
	dataStream := strings.Split(tt.dataStream, "")
	for index := range strings.Split(tt.dataStream, "") {
		sliceStartAt := 0
		sliceEndAt := index

		if index > numberOfDistinctChars {
			sliceStartAt = index - numberOfDistinctChars
		}

		lastFour := dataStream[sliceStartAt:sliceEndAt]
		if isUnique(lastFour, numberOfDistinctChars) {
			return index
		}
	}
	return 0
}

func isUnique(s []string, numberOfDistinctChars int) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && i != j {
				return false
			}
		}
	}
	return len(s) == numberOfDistinctChars
}
