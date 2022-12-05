package suplystack_test

import (
	"testing"

	suplystack "github.com/mhdiiilham/adventofcode/2022/suply-stack"
	"github.com/stretchr/testify/assert"
)

const input = "../../input/2022/suply_stack_test.txt"

func TestLoadInput(t *testing.T) {
	assertion := assert.New(t)
	expectedStacks := map[int][]string{
		1: {"N", "Z"},
		2: {"D", "C", "M"},
		3: {"P"},
	}

	suplyStack := suplystack.New(input)
	errLoadInput := suplyStack.LoadInput()
	assertion.NoError(errLoadInput)
	assertion.Equal(expectedStacks, suplyStack.StackOne)
	assertion.Equal(expectedStacks, suplyStack.StackTwo)
}

func TestCrateMover9000(t *testing.T) {
	assertion := assert.New(t)
	expected := "CMZ"

	suplyStack := suplystack.New(input)
	errLoadInput := suplyStack.LoadInput()
	assertion.NoError(errLoadInput)

	actual := suplyStack.CrateMover9000()
	assertion.Equal(expected, actual)
}

func TestCrateMover9001(t *testing.T) {
	assertion := assert.New(t)
	expected := "MCD"

	suplyStack := suplystack.New(input)
	errLoadInput := suplyStack.LoadInput()
	assertion.NoError(errLoadInput)

	actual := suplyStack.CrateMover9001()
	assertion.Equal(expected, actual)
}
