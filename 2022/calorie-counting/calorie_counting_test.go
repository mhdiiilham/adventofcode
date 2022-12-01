package caloriecounting_test

import (
	"testing"

	caloriecounting "github.com/mhdiiilham/adventofcode/2022/calorie-counting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type calorieCountingTestSuite struct {
	suite.Suite
	calorieCounting *caloriecounting.CalorieCounting
}

func TestCalorieCounting(t *testing.T) {
	suite.Run(t, new(calorieCountingTestSuite))
}

func (suite *calorieCountingTestSuite) TestGetMostCalories() {
	t := suite.T()

	t.Run("success", func(t *testing.T) {
		assertion := assert.New(t)

		err := suite.calorieCounting.LoadCalories()
		assertion.NoError(err)
		mostCalories := suite.calorieCounting.GetMostCalories()
		assertion.Equal(24000, mostCalories)
	})
}

func (suite *calorieCountingTestSuite) SetupTest() {
	sourcePath := "../../input/2022/elf_calories_test.txt"
	suite.calorieCounting = caloriecounting.NewCalorieCounting(sourcePath)
}
