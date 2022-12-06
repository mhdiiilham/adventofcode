package tuningtrouble

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type TuningTrouble struct {
	InputSource string
	Datastream  string
}

const (
	packet  = 4
	message = 14
)

func New(inputSource string) *TuningTrouble {
	return &TuningTrouble{InputSource: inputSource}
}

func (tt *TuningTrouble) LoadInput() (err error) {
	b, err := ioutil.ReadFile(tt.InputSource)
	if err != nil {
		return fmt.Errorf("failed to read input file %s", tt.InputSource)
	}

	tt.Datastream = string(b)
	return nil
}

func (tt *TuningTrouble) GetHowManyCharactersNeedToBeProcessedBeforeTheFirstStartOfPacketMarkerDetected() int {
	return tt.getHowManyCharactersNeedToBeProcessedBeforeTheFirstStart(packet)
}

func (tt *TuningTrouble) GetHowManyCharactersNeedToBeProcessedBeforeTheFirstStartOfMessageDetected() int {
	return tt.getHowManyCharactersNeedToBeProcessedBeforeTheFirstStart(message)
}

func (tt *TuningTrouble) getHowManyCharactersNeedToBeProcessedBeforeTheFirstStart(numberOfDistinctChars int) int {
	dataStream := strings.Split(tt.Datastream, "")
	for index := range strings.Split(tt.Datastream, "") {
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
