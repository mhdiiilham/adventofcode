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
	return tt.getHowManyCharactersNeedToBeProcessedBeforeTheFirstStart(4)
}

func (tt *TuningTrouble) GetHowManyCharactersNeedToBeProcessedBeforeTheFirstStartOfMessageDetected() int {
	return tt.getHowManyCharactersNeedToBeProcessedBeforeTheFirstStart(14)
}

func (tt *TuningTrouble) getHowManyCharactersNeedToBeProcessedBeforeTheFirstStart(n int) int {
	new := []string{}
	for i, char := range strings.Split(tt.Datastream, "") {
		new = append(new, char)
		start := 0
		end := i
		if i > n {
			start = i - n
		}

		copied := copySlice(new)
		lastFour := copied[start:end]
		if unique(lastFour, n) {
			return i
		}

	}
	return 0

}

func copySlice(s []string) []string {
	var result []string
	result = append(result, s...)
	return result
}

func unique(s []string, n int) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && i != j {
				return false
			}
		}
	}

	return len(s) == n
}
