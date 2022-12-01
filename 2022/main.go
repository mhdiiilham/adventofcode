package main

import (
	"fmt"
	"log"

	caloriecounting "github.com/mhdiiilham/adventofcode/2022/calorie-counting"
)

func main() {
	fmt.Println("Advance of Code 2022")
	fmt.Println("--------------------")

	fmt.Println("--Counting Calorie--")
	calorieCounting := caloriecounting.NewCalorieCounting("input/2022/elf_calories.txt")
	err := calorieCounting.LoadCalories()
	if err != nil {
		log.Fatalf("failed load calories: %v", err)
	}
	mostCalories := calorieCounting.GetMostCalories()
	fmt.Printf("Most total calories: %d\n", mostCalories)

}
