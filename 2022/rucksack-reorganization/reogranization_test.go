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

func TestFailedReadInput(t *testing.T) {
	err := rucksackreorganization.NewRucksack("").LoadInput()
	assert.Error(t, err)
}

func BenchmarkGetSumOfThePrioritiesOfItems(b *testing.B) {
	rucksack := rucksackreorganization.NewRucksack(input)
	if err := rucksack.LoadInput(); err != nil {
		b.FailNow()
	}

	for i := 0; i < b.N; i++ {
		rucksack.GetSumOfThePrioritiesOfItems()
	}
}

func BenchmarkGetSumOfBadgePrioritiesItem(b *testing.B) {
	rucksack := rucksackreorganization.NewRucksack(input)
	if err := rucksack.LoadInput(); err != nil {
		b.FailNow()
	}

	for i := 0; i < b.N; i++ {
		rucksack.GetSumOfBadgePrioritiesItem()
	}
}
