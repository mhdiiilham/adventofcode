package input

import (
	"io/ioutil"
	"strings"
)

type InputSeparator string

var (
	OneEnter InputSeparator = "\n"
	TwoEnter InputSeparator = "\n\n"
)

func ReadByLines(inputSource string, seperators ...InputSeparator) ([]string, error) {
	var result []string

	var seperator InputSeparator = "\n"
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
