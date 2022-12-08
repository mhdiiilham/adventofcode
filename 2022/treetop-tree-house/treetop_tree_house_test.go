package treetoptreehouse_test

import (
	"testing"

	treetoptreehouse "github.com/mhdiiilham/adventofcode/2022/treetop-tree-house"
	"github.com/stretchr/testify/assert"
)

const input = "../../input/2022/treetop_tree_house_test.txt"

func TestTreeTopTreeHouse(t *testing.T) {
	assertion := assert.New(t)
	expected := 21

	c, err := treetoptreehouse.NewClient(input)
	assertion.NoError(err)

	numberOfTreeVisibleFromOutsideGrid := c.GetHowManyVisibleTreesAvailableFromOutsideGrid()
	assertion.Equal(expected, numberOfTreeVisibleFromOutsideGrid, "number of tree that visible from outside grid should be %d", expected)
}

func TestGetHighestSchenicScores(t *testing.T) {
	assertion := assert.New(t)
	expected := 8

	c, err := treetoptreehouse.NewClient(input)
	assertion.NoError(err)

	c.GetHowManyVisibleTreesAvailableFromOutsideGrid()
	highestScenicScore := c.GetHigestScenicScores()
	assertion.Equal(expected, highestScenicScore, "the highest scenic score possible for any tree should be %d", expected)

}

func TestFailedReadInput(t *testing.T) {
	assertion := assert.New(t)

	c, err := treetoptreehouse.NewClient("")
	assertion.Error(err)
	assertion.Nil(c)
}
