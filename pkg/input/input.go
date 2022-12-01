package input

import (
	"io/ioutil"
	"strings"
)

type InputSeperator string

var (
	OneEnter InputSeperator = "\n"
	TwoEnter InputSeperator = "\n\n"
)

func ReadByLines(inputSource string, seperators ...InputSeperator) ([]string, error) {
	var result []string

	var seperator InputSeperator = "\n"
	if len(seperators) == 1 {
		seperator = seperators[0]
	}

	source, err := ioutil.ReadFile(inputSource)
	if err != nil {
		return nil, err
	}

	result = append(result, strings.Split(string(source), string(seperator))...)

	return result, nil
}
