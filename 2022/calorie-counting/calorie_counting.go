package caloriecounting

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type CalorieCounting struct {
	InputSource  string
	Calories     [][]int
	SumCalories  []int
	MostCalories int
}

func NewCalorieCounting(inputSource string) *CalorieCounting {
	return &CalorieCounting{InputSource: inputSource}
}

func (c *CalorieCounting) GetMostCalories() (mostCalories int) {
	for _, groupCalories := range c.Calories {
		tempValue := 0
		for _, group := range groupCalories {
			tempValue += group
		}

		if tempValue > c.MostCalories {
			c.MostCalories = tempValue
		}
	}

	return c.MostCalories
}

func (c *CalorieCounting) GetTotalTopThreeCaloriesCarriedByElves() int {
	c.SumAllCaloriesByElves().SortSumCaloriesByLargest()
	topThree := c.SumCalories[:3]
	sum := 0
	for _, v := range topThree {
		sum += v
	}

	return sum
}

func (c *CalorieCounting) SortSumCaloriesByLargest() *CalorieCounting {
	sort.Sort(sort.Reverse(sort.IntSlice(c.SumCalories)))
	return c
}

func (c *CalorieCounting) SumAllCaloriesByElves() *CalorieCounting {
	for _, groupCalories := range c.Calories {
		totalCalories := 0
		for _, group := range groupCalories {
			totalCalories += group
		}
		c.SumCalories = append(c.SumCalories, totalCalories)
	}

	return c
}

func (c *CalorieCounting) LoadCalories() (err error) {
	b, err := ioutil.ReadFile(c.InputSource)
	if err != nil {
		return fmt.Errorf("failed to read file: %s", c.InputSource)
	}

	result := strings.Split(string(b), "\n\n")
	c.GroupingInput(result)

	return nil
}

func (c *CalorieCounting) GroupingInput(i []string) {
	for _, row := range i {
		elfCalorie := []int{}
		calories := strings.Split(row, "\n")
		for _, calorie := range calories {
			calorieInt := 0
			calorieInt, _ = strconv.Atoi(calorie)
			elfCalorie = append(elfCalorie, calorieInt)
		}
		c.Calories = append(c.Calories, elfCalorie)
	}
}
