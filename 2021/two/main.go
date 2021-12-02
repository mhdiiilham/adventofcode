package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/mhdiiilham/adventofcode/pkg/reader"
)

type Navigation struct {
	Action string
	Value  int
}

func main() {
	fmt.Println("advent of code day: 2")
	file := flag.String("file", "", "to read advent of puzzle input")
	flag.Parse()

	r := reader.NewReader(*file)
	err := r.ReadInput()
	if err != nil {
		panic(err)
	}

	inputs := r.RawResult
	parsedInput := ParseInput(inputs)

	fmt.Printf("Result of multiplying horizontal and vertical position: %d\n", CalculatePosition(parsedInput))
}

func CalculatePosition(navigations []Navigation) (multiplyResult int) {
	horizontal := 0
	vertical := 0

	for _, nav := range navigations {
		switch nav.Action {
		case "forward":
			horizontal = horizontal + nav.Value
		case "down":
			vertical = vertical + nav.Value
		case "up":
			vertical = vertical - nav.Value
		}
	}

	return horizontal * vertical
}

func ParseInput(rawInputs []string) (navigations []Navigation) {
	for _, input := range rawInputs {
		rawInput := strings.Split(input, " ")
		value, _ := strconv.Atoi(rawInput[1])
		navigations = append(navigations, Navigation{
			Action: rawInput[0],
			Value:  value,
		})
	}
	return
}
