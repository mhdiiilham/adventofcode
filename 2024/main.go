package main

import (
	"fmt"

	historianhysteria "github.com/mhdiiilham/adventofcode/2024/1_historian-hysteria"
	rednosereport "github.com/mhdiiilham/adventofcode/2024/2_red-nose-report"
	mullitover "github.com/mhdiiilham/adventofcode/2024/3_mull_it_over"
)

func main() {
	fmt.Println("Advent of code 2024")

	if dayOne, err := historianhysteria.NewSolver("input/2024/1.txt"); err == nil {
		partOne, partTwo := dayOne.Solve()
		fmt.Printf("[Historian Hysteria]\nThe total distance between your lists is: %d.\nTheir similarity score is: %d.\n\n", partOne, partTwo)
	}

	if dayTwo, err := rednosereport.NewSolver("input/2024/2.txt"); err == nil {
		partOne, partTwo := dayTwo.Solve()
		fmt.Printf("[Red-Nose Reports]\nNumber of safe reports are: %d.\nNumber of safe reports after removing one any element are: %d.\n\n", partOne, partTwo)
	}

	if dayThree, err := mullitover.NewSolver("input/2024/3.txt"); err == nil {
		partOne, partTwo := dayThree.Solve()
		fmt.Printf("[Mull It Over]\nResults of the multiplications: %d.\nTotal of enabled multiplications: %d.\n\n.", partOne, partTwo)
	}
}
