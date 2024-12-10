package main

import (
	"fmt"

	historianhysteria "github.com/mhdiiilham/adventofcode/2024/1_historian-hysteria"
)

func main() {
	fmt.Println("Advent of code 2024")

	if dayOne, err := historianhysteria.NewSolver("input/2024/1.txt"); err == nil {
		partOne, partTwo := dayOne.Solve()
		fmt.Printf("[Historian Hysteria]\nThe total distance between your lists is: %d.\nTheir similarity score is: %d.\n", partOne, partTwo)
	}
}
