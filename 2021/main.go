package main

import (
	"fmt"
	"log"

	"github.com/mhdiiilham/adventofcode/2021/one"
	"github.com/mhdiiilham/adventofcode/2021/two"
	"github.com/mhdiiilham/adventofcode/pkg/reader"
)

func main() {
	fmt.Println("===== advent of code 2021 =====")

	fmt.Println("=== DAY: 1 ===")
	r := reader.NewReader("input/2021/one.txt")
	err := r.ReadInput()
	if err != nil {
		log.Printf("[ERROR] - reading day one puzzle input: %v\n", err)
	}
	puzzleOneInput, err := r.GetResultAsSliceOfInt()
	if err != nil {
		log.Printf("[ERROR] - reading day one puzzle input GetResultAsSliceOfInt: %v\n", err)
	}
	sonarSweep := one.NewSonarSweep(puzzleOneInput)
	fmt.Printf("Number of deepth measurement increased: %d\n", sonarSweep.NumberOfMeasurementIncreased(sonarSweep.Measurements))
	fmt.Printf("Number of three-measurements windows: %d\n\n", sonarSweep.ThreeWindowMeasurements())

	fmt.Println("=== DAY: 2 ===")
	readerTwo := r.SetFileInput("input/2021/two.txt")
	if err := readerTwo.ReadInput(); err != nil {
		log.Printf("[ERROR] - reading day two puzzle input: %v\n", err)
	}
	dive := two.NewDive(readerTwo.RawResult)
	fmt.Printf("Result of multiplying horizontal and vertical position: %d\n", dive.CalculatePosition())
	fmt.Printf("Result of multiplying horizontal and vertical position (Part 2): %d\n\n", dive.CalculatePositionPart2())

	fmt.Println("===== end of advent of code 2021 =====")
}
