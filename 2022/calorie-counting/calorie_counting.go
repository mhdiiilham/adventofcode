package caloriecounting

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type CalorieCounting struct {
	InputSource  string
	Calories     [][]int
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

func (c *CalorieCounting) LoadCalories() (err error) {
	b, err := ioutil.ReadFile(c.InputSource)
	if err != nil {
		return err
	}

	result := strings.Split(string(b), "\n\n")
	if err := c.GroupingInput(result); err != nil {
		return err
	}

	return nil
}

func (c *CalorieCounting) GroupingInput(i []string) (err error) {
	for _, row := range i {
		elfCalorie := []int{}
		calories := strings.Split(row, "\n")
		for _, calorie := range calories {
			calorieInt, err := strconv.Atoi(calorie)
			if err != nil {
				elfCalorie = append(elfCalorie, 0)
			}
			elfCalorie = append(elfCalorie, calorieInt)
		}
		c.Calories = append(c.Calories, elfCalorie)
	}

	return nil
}
