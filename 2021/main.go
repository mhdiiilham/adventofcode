package main

import (
	"fmt"
	"log"

	"github.com/mhdiiilham/adventofcode/2021/binarydiagnostic"
	"github.com/mhdiiilham/adventofcode/2021/dive"
	"github.com/mhdiiilham/adventofcode/2021/sonarsweep"
	"github.com/mhdiiilham/adventofcode/pkg/reader"
)

func main() {
	fmt.Printf("===== [advent of code 2021] =====\n\n")

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
	sonarSweep := sonarsweep.NewSonarSweep(puzzleOneInput)
	fmt.Printf("Number of deepth measurement increased: %d\n", sonarSweep.NumberOfMeasurementIncreased(sonarSweep.Measurements))
	fmt.Printf("Number of three-measurements windows: %d\n\n", sonarSweep.ThreeWindowMeasurements())

	fmt.Println("=== DAY: 2 ===")
	readerTwo := r.SetFileInput("input/2021/two.txt")
	if err := readerTwo.ReadInput(); err != nil {
		log.Printf("[ERROR] - reading day two puzzle input: %v\n", err)
	}
	Dive := dive.NewDive(readerTwo.RawResult)
	fmt.Printf("Result of multiplying horizontal and vertical position: %d\n", Dive.CalculatePosition())
	fmt.Printf("Result of multiplying horizontal and vertical position (Part 2): %d\n\n", Dive.CalculatePositionPart2())

	fmt.Println("=== Day: 3 ===")
	readerThree := r.SetFileInput("input/2021/three.txt")
	if err := readerThree.ReadInput(); err != nil {
		log.Printf("[ERROR] - reading day three puzzle input: %v\n", err)
	}

	diagnoster := binarydiagnostic.NewDiagnoster(readerThree.RawResult)
	fmt.Printf("power consumption of the submarine: %d\n", diagnoster.CalculatePowerConsumption())
	fmt.Printf("submarine life support rating: %d\n\n", diagnoster.CalculateLifeSupportRating())

	fmt.Println("===== [end of advent of code 2021] =====")
}
