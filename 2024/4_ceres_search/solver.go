package ceressearch

import (
	"os"
	"strings"
)

// Directions: (row offset, col offset)
var directions = [8][2]int{
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 0},   // Down
	{-1, 0},  // Up
	{1, 1},   // Down-right
	{1, -1},  // Down-left
	{-1, 1},  // Up-right
	{-1, -1}, // Up-left
}

// Solver struct hold necessary data needed to solve the Ceres Search challenge.
type Solver struct {
	inputFilePath string
	grid          [][]string
}

// NewSolver function return the new solver instance of challenge Ceres Search.
func NewSolver(inputFilePath string) (*Solver, error) {
	s := &Solver{inputFilePath: inputFilePath}
	return s.read()
}

// Solve function solve the Ceres Search challenge.
func (s *Solver) Solve() (int, int) {
	return s.partOne(), s.partTwo()
}

func (s *Solver) partOne() int {
	return s.countWordOccurrences(s.grid, "XMAS")
}

func (s *Solver) partTwo() int {
	return s.countXMASOccurrences(s.grid)
}

func (s *Solver) findWord(grid [][]string, word string, startRow, startCol, dirRow, dirCol int) bool {
	wordLen := len(word)
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < wordLen; i++ {
		r := startRow + i*dirRow
		c := startCol + i*dirCol

		// Check boundaries
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}

		// Check character match
		if grid[r][c] != string(word[i]) {
			return false
		}
	}
	return true
}

func (s *Solver) countWordOccurrences(grid [][]string, word string) int {
	count := 0
	rows, cols := len(grid), len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Check in all 8 directions
			for _, dir := range directions {
				if s.findWord(grid, word, row, col, dir[0], dir[1]) {
					count++
				}
			}
		}
	}
	return count
}

func (s *Solver) findXMAS(grid [][]string, row, col int) bool {
	rows, cols := len(grid), len(grid[0])

	// Check bounds for the arms of the X
	if row-1 < 0 || row+1 >= rows || col-1 < 0 || col+1 >= cols {
		return false
	}

	// Top-left to bottom-right arm
	topLeft := (grid[row-1][col-1] == "M" && grid[row][col] == "A" && grid[row+1][col+1] == "S") ||
		(grid[row-1][col-1] == "S" && grid[row][col] == "A" && grid[row+1][col+1] == "M")

	// Top-right to bottom-left arm
	topRight := (grid[row-1][col+1] == "M" && grid[row][col] == "A" && grid[row+1][col-1] == "S") ||
		(grid[row-1][col+1] == "S" && grid[row][col] == "A" && grid[row+1][col-1] == "M")

	// Both arms must form the X-MAS pattern
	return topLeft && topRight
}

func (s *Solver) countXMASOccurrences(grid [][]string) int {
	count := 0
	rows, cols := len(grid), len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Look for X-MAS centered at (row, col)
			if grid[row][col] == "A" && s.findXMAS(grid, row, col) {
				count++
			}
		}
	}
	return count
}

func (s *Solver) read() (*Solver, error) {
	inputInBytes, err := os.ReadFile(s.inputFilePath)
	if err != nil {
		return s, err
	}

	var grid [][]string
	for _, horizontal := range strings.Split(string(inputInBytes), "\n") {
		grid = append(grid, strings.Split(horizontal, ""))
	}

	s.grid = grid
	return s, nil
}
