package main

import (
	"flag"
	"fmt"

	"github.com/mhdiiilham/adventofcode/pkg/reader"
)

func main() {
	fmt.Println("advent of code day: 1")
	file := flag.String("file", "", "to read advent of puzzle input")
	flag.Parse()

	r := reader.NewReader(*file)
	err := r.ReadInput()
	if err != nil {
		panic(err)
	}

	inputs, err := r.GetResultAsSliceOfInt()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of deepth measurement increased: %d\n", NumberOfMeasurementIncreased(inputs))
	fmt.Printf("Number of three-measurements windows: %d\n", ThreeWindowMeasurements(inputs))
}

func NumberOfMeasurementIncreased(measurements []int) (increasedNTime int) {
	if len(measurements) == 0 {
		return 0
	}

	prevMeasurement := measurements[0]
	for _, measurement := range measurements {
		if measurement > prevMeasurement {
			increasedNTime++
		}
		prevMeasurement = measurement
	}

	return
}

func ThreeWindowMeasurements(measurements []int) (increasedNTime int) {
	threeMeasurementsResult := []int{}
	lastThreeIndex := 0

	for lastThreeIndex+2 < len(measurements) {
		firstIndex := lastThreeIndex
		secondIndex := lastThreeIndex + 1
		thirdIndex := lastThreeIndex + 2

		// Sum of three-measurements windows
		sum := measurements[firstIndex] + measurements[secondIndex] + measurements[thirdIndex]
		threeMeasurementsResult = append(threeMeasurementsResult, sum)

		lastThreeIndex++
	}

	return NumberOfMeasurementIncreased(threeMeasurementsResult)
}
