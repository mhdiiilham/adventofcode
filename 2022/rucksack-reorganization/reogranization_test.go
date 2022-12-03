package rucksackreorganization_test

import (
	"testing"

	rucksackreorganization "github.com/mhdiiilham/adventofcode/2022/rucksack-reorganization"
	"github.com/stretchr/testify/assert"
)

const input = "../../input/2022/rucksack_reorganization_test.txt"

func TestSumOfThePrioritiesOfItems(t *testing.T) {
	assertion := assert.New(t)
	expectedSum := 157

	rucksack := rucksackreorganization.NewRucksack(input)
	errLoadInput := rucksack.LoadInput()
	assertion.NoError(errLoadInput)

	gotSum := rucksack.GetSumOfThePrioritiesOfItems()
	assertion.Equal(expectedSum, gotSum)
}

func TestGetSumOfBadgePrioritiesItem(t *testing.T) {
	assertion := assert.New(t)
	rucksack := rucksackreorganization.NewRucksack(input)
	errLoadInput := rucksack.LoadInput()
	assertion.NoError(errLoadInput)

	expectedSum := 70
	actualSum := rucksack.GetSumOfBadgePrioritiesItem()
	assertion.Equal(expectedSum, actualSum)

}
