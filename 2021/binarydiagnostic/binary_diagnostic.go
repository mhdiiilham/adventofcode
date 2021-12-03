package binarydiagnostic

import (
	"log"
	"strconv"
	"strings"
)

type DiagnosticReport struct {
	FirstMostCommon  int
	SecondMostCommon int
	ThirdMostCommon  int
	FourthMostCommon int
	FifthMostCommon  int
}

type diagnoster struct {
	rawInput []string
	DiagnosticReport
}

func NewDiagnoster(rawInput []string) *diagnoster {
	return &diagnoster{
		rawInput: rawInput,
	}
}

func (d *diagnoster) CalculatePowerConsumption() int64 {
	gamma := ""
	epsilon := ""

	for i := range strings.Split(d.rawInput[0], "") {
		nol := 0
		one := 0

		for _, report := range d.rawInput {
			reports := strings.Split(report, "")
			if reports[i] == "0" {
				nol++
			}

			if reports[i] == "1" {
				one++
			}
		}

		if nol > one {
			gamma += "0"
			epsilon += "1"
		}

		if one > nol {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaValue, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Printf("[ERROR] parsing gamma binary to int: %v\n", err)
		return 0
	}

	epsilonValue, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Printf("[ERROR] parsing gamma binary to int: %v\n", err)
		return 0
	}

	powerComsumption := gammaValue * epsilonValue
	return powerComsumption
}

func (d *diagnoster) CalculateLifeSupportRating() (rating int64) {
	oxygenGeneratorRatingBits := d.getOxygenGeneratorValue()
	oxygenGeneratorRating, err := strconv.ParseInt(oxygenGeneratorRatingBits, 2, 64)
	if err != nil {
		log.Printf("[ERROR] parsing oxygenGeneratorRatingBits to int: %v\n", err)
	}

	getCO2ScrubberValue := d.getCO2ScrubberValue()
	CO2ScrubberRating, err := strconv.ParseInt(getCO2ScrubberValue, 2, 64)
	if err != nil {
		log.Printf("[ERROR] parsing oxygenGeneratorRatingBits to int: %v\n", err)
	}

	return oxygenGeneratorRating * CO2ScrubberRating
}

func (d *diagnoster) getOxygenGeneratorValue() (bit string) {
	tempBits := d.rawInput
	maxPos := len(tempBits[0]) - 1

	for len(tempBits) != 1 {
		currentPos := 0
		for currentPos <= maxPos {
			nol := []string{}
			one := []string{}

			for _, report := range tempBits {
				b := strings.Split(report, "")

				if b[currentPos] == "1" {
					one = append(one, report)
				}

				if b[currentPos] == "0" {
					nol = append(nol, report)
				}
			}

			if len(nol) > len(one) {
				tempBits = nol
			} else {
				tempBits = one
			}
			currentPos++
			if len(tempBits) == 1 {
				break
			}
		}

		if len(tempBits) == 1 {
			break
		}
	}

	return tempBits[0]
}

func (d *diagnoster) getCO2ScrubberValue() (bit string) {
	tempBits := d.rawInput
	maxPos := len(tempBits[0]) - 1

	for len(tempBits) != 1 && len(tempBits) > 0 {
		currentPos := 0
		for currentPos <= maxPos {
			nol := []string{}
			one := []string{}

			for _, report := range tempBits {
				b := strings.Split(report, "")

				if b[currentPos] == "1" {
					one = append(one, report)
				}

				if b[currentPos] == "0" {
					nol = append(nol, report)
				}
			}

			if len(nol) > len(one) {
				tempBits = one
			} else {
				tempBits = nol
			}

			currentPos++
			if len(tempBits) == 1 {
				break
			}
		}

		if len(tempBits) == 1 {
			break
		}

	}

	return tempBits[0]
}
