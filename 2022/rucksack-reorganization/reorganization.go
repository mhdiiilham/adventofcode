package rucksackreorganization

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mhdiiilham/adventofcode/pkg/input"
)

type Rucksack struct {
	InputSource             string
	Input                   []string
	RuckSack                [][]string
	ContainedInBothRucksack []string
	ElvesGroup              [][]string
}

func NewRucksack(inputSource string) *Rucksack {
	return &Rucksack{InputSource: inputSource}
}

func (r *Rucksack) LoadInput() (err error) {
	r.Input, err = input.ReadByLines(r.InputSource, input.OneEnter)
	if err != nil {
		return err
	}

	for _, row := range r.Input {
		inputRow := strings.Split(row, "")

		half := len(inputRow) / 2
		firstHalf := inputRow[:half]
		secondHalf := inputRow[half:]

		rucksackRow := []string{strings.Join(firstHalf, ""), strings.Join(secondHalf, "")}
		r.RuckSack = append(r.RuckSack, rucksackRow)
	}

	return nil
}

func (r *Rucksack) GetSumOfThePrioritiesOfItems() int {
	for _, row := range r.RuckSack {

		firstHalf := row[0]
		secondHalf := row[1]

		pattern := fmt.Sprintf("[(%s)]", firstHalf)
		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(secondHalf, 1)
		for _, match := range matches {
			r.ContainedInBothRucksack = append(r.ContainedInBothRucksack, match...)
		}
	}
	result := 0
	for _, item := range r.ContainedInBothRucksack {
		result += charToNumber(item)
	}
	return result
}

func charToNumber(char string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, item := range strings.Split(alphabet, "") {
		if item == char {
			return i + 1
		}
	}

	return 0
}

func (r *Rucksack) GetSumOfBadgePrioritiesItem() int {
	r.groupingElves()

	found := []string{}
	for i := 0; i < len(r.ElvesGroup); i++ {
		matchesGroup := []string{}
		firstGroup := r.ElvesGroup[i][0]
		secondGroup := r.ElvesGroup[i][1]
		thirdGroup := r.ElvesGroup[i][2]

		pattern := fmt.Sprintf("[(%s)]", firstGroup)
		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(secondGroup, -1)
		for _, match := range matches {
			matchesGroup = append(matchesGroup, match[0])
		}

		matchAll := []string{}
		secondPattern := fmt.Sprintf("[(%s)]", strings.Join(matchesGroup, ""))
		re = regexp.MustCompile(secondPattern)
		matches = re.FindAllStringSubmatch(thirdGroup, -1)
		for _, match := range matches {
			matchAll = append(matchAll, match[0])
		}

		found = append(found, matchAll[0])
	}

	result := 0
	for _, f := range found {
		result += charToNumber(f)
	}

	return result
}

func (r *Rucksack) groupingElves() {
	count := 1
	group := []string{}

	for _, row := range r.Input {
		group = append(group, row)
		count++
		if count > 3 {
			count = 1
			r.ElvesGroup = append(r.ElvesGroup, group)
			group = []string{}
		}
	}
}
