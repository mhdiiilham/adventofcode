package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	caloriecounting "github.com/mhdiiilham/adventofcode/2022/calorie-counting"
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

}
