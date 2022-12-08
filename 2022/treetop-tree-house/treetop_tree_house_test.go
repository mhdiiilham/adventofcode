package treetoptreehouse_test

import (
	"testing"

	treetoptreehouse "github.com/mhdiiilham/adventofcode/2022/treetop-tree-house"
	"github.com/stretchr/testify/assert"
)

const input = "../../input/2022/treetop_tree_house_test.txt"

func TestGetHowManyVisibleTreesAvailableFromOutsideGrid(t *testing.T) {
	assertion := assert.New(t)
	expected := 21
	expected2 := 8

	c, err := treetoptreehouse.NewClient(input)
	assertion.NoError(err)

	part1, part2 := c.GetHowManyVisibleTreesAvailableFromOutsideGrid()
	assertion.Equal(expected, part1, "number of tree that available should be %d", expected)
	assertion.Equal(expected2, part2, "number of tree that available should be %d", expected2)

}
