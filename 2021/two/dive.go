package two

import (
	"strconv"
	"strings"
)

type (
	Navigation struct {
		Action string
		Value  int
	}

	dive struct {
		Navigations []Navigation
	}
)

func NewDive(rawPuzzleInput []string) *dive {
	navigations := ParseInput(rawPuzzleInput)
	return &dive{Navigations: navigations}
}

func (d *dive) CalculatePosition() (multiplyResult int) {
	horizontal := 0
	vertical := 0

	for _, nav := range d.Navigations {
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

func (d *dive) CalculatePositionPart2() (multiplyResult int) {
	horizontal := 0
	vertical := 0
	aim := 0

	for _, nav := range d.Navigations {
		switch nav.Action {
		case "forward":
			horizontal = horizontal + nav.Value
			if aim > 0 {
				aim := nav.Value * aim
				vertical += aim
			}
		case "down":
			aim = aim + nav.Value
		case "up":
			aim = aim - nav.Value
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
