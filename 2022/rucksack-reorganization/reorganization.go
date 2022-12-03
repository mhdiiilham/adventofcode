package rucksackreorganization

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mhdiiilham/adventofcode/pkg/input"
)

type Rucksack struct {
	InputSource string
	Input       []string
	RuckSack    [][]string
	ElvesGroup  [][]string
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
	containerCh := make(chan string, len(r.RuckSack))

	for _, row := range r.RuckSack {
		go r.findMatchingInOtherHalf(row, containerCh)
	}

	result := 0
	for i := 0; i < len(r.RuckSack); i++ {
		char := <-containerCh
		result += charToNumber(char)
	}
	return result
}

func (r *Rucksack) findMatchingInOtherHalf(row []string, contaier chan<- string) {
	firstHalf := row[0]
	secondHalf := row[1]

	pattern := fmt.Sprintf("[(%s)]", firstHalf)
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(secondHalf, -1)
	contaier <- matches[0]
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

	foundCh := make(chan string, len(r.ElvesGroup))
	for i := 0; i < len(r.ElvesGroup); i++ {
		go r.findMatchesInGroup(i, foundCh)
	}

	result := 0
	for i := 0; i < len(r.ElvesGroup); i++ {
		char := <-foundCh
		result += charToNumber(char)
	}

	return result
}

func (r *Rucksack) findMatchesInGroup(i int, matched chan<- string) {
	matchesGroup := []string{}
	firstGroup := r.ElvesGroup[i][0]
	secondGroup := r.ElvesGroup[i][1]
	thirdGroup := r.ElvesGroup[i][2]

	pattern := fmt.Sprintf("[(%s)]", firstGroup)
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(secondGroup, -1)
	matchesGroup = append(matchesGroup, matches...)

	matchAll := []string{}
	secondPattern := fmt.Sprintf("[(%s)]", strings.Join(matchesGroup, ""))
	re = regexp.MustCompile(secondPattern)
	matches = re.FindAllString(thirdGroup, -1)
	matchAll = append(matchAll, matches...)

	found := matchAll[0]
	matched <- found
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
