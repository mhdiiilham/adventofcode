package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	caloriecounting "github.com/mhdiiilham/adventofcode/2022/calorie-counting"
	campcleanup "github.com/mhdiiilham/adventofcode/2022/camp-cleanup"
	nospaceleftondevice "github.com/mhdiiilham/adventofcode/2022/no-space-left-on-device"
	rockpaperscissor "github.com/mhdiiilham/adventofcode/2022/rock-paper-scissor"
	rucksackreorganization "github.com/mhdiiilham/adventofcode/2022/rucksack-reorganization"
	suplystack "github.com/mhdiiilham/adventofcode/2022/suply-stack"
	treetoptreehouse "github.com/mhdiiilham/adventofcode/2022/treetop-tree-house"
	tuningtrouble "github.com/mhdiiilham/adventofcode/2022/tuning-trouble"
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

	color.Green("5. Supply Stacks")
	suplyStack := suplystack.New("input/2022/suply_stack.txt")
	if err := suplyStack.LoadInput(); err != nil {
		log.Fatalf("faile to load suply stack input: %v", err)
	}
	crateMover9000 := suplyStack.CrateMover9000()
	crateMover9001 := suplyStack.CrateMover9001()
	fmt.Printf("After the rearrangement procedure completes (crateMover9000), crate that ends up on top of each stack: %s \n", crateMover9000)
	fmt.Printf("After the rearrangement procedure completes (crateMover9001), crate that ends up on top of each stack: %s \n", crateMover9001)
	color.Blue("--------------------")

	color.Green("6. Tuning Trouble")
	tuningTrouble := tuningtrouble.New("input/2022/tuning_trouble.txt")
	if err := tuningTrouble.LoadInput(); err != nil {
		log.Fatalf("faile to load tuning trouble input: %v", err)
	}
	charMarkAfter := tuningTrouble.GetHowManyCharactersNeedToBeProcessedBeforeTheFirstStartOfPacketMarkerDetected()
	message := tuningTrouble.GetHowManyCharactersNeedToBeProcessedBeforeTheFirstStartOfMessageDetected()
	fmt.Printf("How many characters need to be processed before the first start-of-packet marker is detected: %d\n", charMarkAfter)
	fmt.Printf("How many characters need to be processed before the first start-of-message marker is detected: %d\n", message)
	color.Blue("--------------------")

	color.Green("7. No Space Left On Device")
	noSpaceLeftOnDevice, err := nospaceleftondevice.New("input/2022/nslod.txt")
	if err != nil {
		log.Fatalf("failed to load No Space Left On Device input: %v", err)
	}
	totalSizeOfDirLookDir := noSpaceLeftOnDevice.GetTotalSizeOfDirectories()
	totalSizedOfDirNeededToBeRemoved := noSpaceLeftOnDevice.GetTotalSizeDirectoryNeedToDelete()
	fmt.Printf("the sum of the total sizes of those directories: %d\n", totalSizeOfDirLookDir)
	fmt.Printf("the total size of the directory that needed to be removed: %d\n", totalSizedOfDirNeededToBeRemoved)
	color.Blue("--------------------")

	color.Green("Day 8: Treetop Tree House")
	client, err := treetoptreehouse.NewClient("input/2022/treetop_tree_house.txt")
	if err != nil {
		log.Fatalf("failed to load Treetop Tree House input: %v", err)
	}
	numberOfTreeAvailables := client.GetHowManyVisibleTreesAvailableFromOutsideGrid()
	bestScenicScore := client.GetHigestScenicScores()
	fmt.Printf("total many trees are visible from outside the grid: %d\n", numberOfTreeAvailables)
	fmt.Printf("the highest scenic score possible for any tree: %d\n", bestScenicScore)
	color.Blue("--------------------")
}
