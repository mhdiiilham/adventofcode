package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	caloriecounting "github.com/mhdiiilham/adventofcode/2022/calorie-counting"
	rockpaperscissor "github.com/mhdiiilham/adventofcode/2022/rock-paper-scissor"
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

	color.Green("1. Rock Pape Scissors")
	rps, err := rockpaperscissor.NewRockPaperScissors("input/2022/rock_paper_scissors.txt").LoadInput()
	if err != nil {
		log.Fatalf("failed to load day 2 challenge's input: %v", err)
	}
	totalScore := rps.GetTotalScore()
	fixedTotalScore := rps.CalculateRealInput()
	fmt.Printf("Total Score: %d\n", totalScore)
	fmt.Printf("Fixed Total Score: %d\n", fixedTotalScore)
	color.Blue("--------------------")
}
