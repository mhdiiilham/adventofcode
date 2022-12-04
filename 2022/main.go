package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	caloriecounting "github.com/mhdiiilham/adventofcode/2022/calorie-counting"
	campcleanup "github.com/mhdiiilham/adventofcode/2022/camp-cleanup"
	rockpaperscissor "github.com/mhdiiilham/adventofcode/2022/rock-paper-scissor"
	rucksackreorganization "github.com/mhdiiilham/adventofcode/2022/rucksack-reorganization"
)

func main() {
	color.Blue("--------------------")
	color.Blue("Advance of Code 2022")
	color.Blue("--------------------")

	color.Green("1. Counting Calories")
	calorieCounting := caloriecounting.NewCalorieCounting("input/2022/elf_calories.txt")
	err := calorieCounting.LoadCalories()
	if err != nil {
		log.Fatalf("failed load calories: %v", err)
	}
	mostCalories := calorieCounting.GetMostCalories()
	sumOfTopThreeCaloriesCarriedByTheElves := calorieCounting.GetTotalTopThreeCaloriesCarriedByElves()
	fmt.Printf("Most total calories: %d\n", mostCalories)
	fmt.Printf("Calories are those top three Elves carrying in total: %d\n", sumOfTopThreeCaloriesCarriedByTheElves)
	color.Blue("--------------------")

	color.Green("2. Rock Pape Scissors")
	rps, err := rockpaperscissor.NewRockPaperScissors("input/2022/rock_paper_scissors.txt").LoadInput()
	if err != nil {
		log.Fatalf("failed to load day 2 challenge's input: %v", err)
	}
	totalScore := rps.GetTotalScore()
	fixedTotalScore := rps.CalculateRealInput()
	fmt.Printf("Total Score: %d\n", totalScore)
	fmt.Printf("Fixed Total Score: %d\n", fixedTotalScore)
	color.Blue("--------------------")

	color.Green("3. Rucksack Reorganization")
	rucksack := rucksackreorganization.NewRucksack("input/2022/rucksack_organization.txt")
	if err := rucksack.LoadInput(); err != nil {
		log.Fatalf("failed to load input: %v", err)
	}
	sumOfThePrioritiesItem := rucksack.GetSumOfThePrioritiesOfItems()
	sumOfTheBadgePriorities := rucksack.GetSumOfBadgePrioritiesItem()
	fmt.Printf("the sum of the priorities of those item types: %d\n", sumOfThePrioritiesItem)
	fmt.Printf("the sum of the badge priorities of those item types: %d\n", sumOfTheBadgePriorities)
	color.Blue("--------------------")

	color.Green("4. Camp Cleanup")
	campCleanUp := campcleanup.NewCampCleanUp("input/2022/camp_cleanup.txt")
	if err := campCleanUp.LoadInput(); err != nil {
		log.Fatalf("faile to load camp cleanup input: %v", err)
	}

	numbersOfAssignmentThatFullyContainsTheOtherPairAssignment := campCleanUp.GetNumbersOfAssignmentThatFullyContainsTheOtherPairAssignment()
	numbersOfAssignmentPairsDoTheRangeOverlap := campCleanUp.GetNumbersOfAssignmentPairsDoTheRangeOverlap()
	fmt.Printf("total assignment pairs that does one range fully contain the other: %d\n", numbersOfAssignmentThatFullyContainsTheOtherPairAssignment)
	fmt.Printf("total assignment pairs do the ranges overlap: %d\n", numbersOfAssignmentPairsDoTheRangeOverlap)
	color.Blue("--------------------")
}
