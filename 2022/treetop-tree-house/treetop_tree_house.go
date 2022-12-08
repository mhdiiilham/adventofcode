package treetoptreehouse

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mhdiiilham/adventofcode/pkg/input"
)

type Client struct {
	inputSource string
	treeMap     [][]string
}

func NewClient(inputSource string) (*Client, error) {
	c := &Client{inputSource: inputSource}

	rows, err := input.ReadByLines(c.inputSource, input.OneEnter)
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		c.treeMap = append(c.treeMap, strings.Split(row, ""))
	}

	return c, nil
}

func (c *Client) GetHowManyVisibleTreesAvailableFromOutsideGrid() (int, int) {
	edges := 0
	count := 0
	bestViews := map[string][]int{}

	for i := 0; i < len(c.treeMap); i++ {
		for j := 0; j < len(c.treeMap[i]); j++ {
			key := fmt.Sprintf("%d-%d", i, j)
			if i == 0 {
				edges++
			} else if j == 0 && i != 0 {
				edges++
			} else if i == len(c.treeMap)-1 && j != 0 {
				edges++
			} else if j == len(c.treeMap[i])-1 {
				edges++
			} else {
				current, _ := strconv.Atoi(c.treeMap[i][j])
				edgesOnTop := true
				for m := i - 1; m >= 0; m-- {
					t, _ := strconv.Atoi(c.treeMap[m][j])
					if t >= current {
						edgesOnTop = false
						break
					}
				}

				edgesToBottom := true
				for m := i + 1; m < len(c.treeMap); m++ {
					t, _ := strconv.Atoi(c.treeMap[m][j])
					if t >= current {
						edgesToBottom = false
						break
					}
				}

				edgesToLeft := true
				treeToTheLeft := c.treeMap[i][:j]
				for _, m := range treeToTheLeft {
					t, _ := strconv.Atoi(m)
					if t >= current {
						edgesToLeft = false
						break
					}
				}

				edgesToRight := true
				treeToTheRight := c.treeMap[i][j+1:]
				for _, m := range treeToTheRight {
					t, _ := strconv.Atoi(m)
					if t >= current {
						edgesToRight = false
						break
					}
				}

				if edgesOnTop || edgesToBottom || edgesToLeft || edgesToRight {
					count++
				}

				left := countToLeft(c.treeMap[i], j)
				right := countToRight(c.treeMap[i], j)
				up := countUp(c.treeMap, i, j)
				down := countDown(c.treeMap, i, j)
				bestViews[key] = append(bestViews[key], left, right, up, down)
			}
		}
	}

	fmt.Println(bestViews)
	max := 0
	for _, views := range bestViews {
		temp := 0
		for _, view := range views {
			if view != 0 {
				if temp == 0 {
					temp = view
				} else {
					temp = temp * view
				}
			}
		}
		if max < temp {
			max = temp
		}
	}
	fmt.Println(max)
	return edges + count, max
}

func countToLeft(trees []string, start int) int {
	result := 0

	for i := start - 1; i >= 0; i-- {
		result++
		n, _ := strconv.Atoi(trees[i])
		tt, _ := strconv.Atoi(trees[start])
		if n >= tt {
			break
		}
	}

	return result
}

func countUp(trees [][]string, current, idx int) int {
	result := 0
	for i := current - 1; i >= 0; i-- {
		result++
		t, _ := strconv.Atoi(trees[current][idx])
		n, _ := strconv.Atoi(trees[i][idx])
		if n >= t {
			break
		}
	}
	return result
}

func countDown(trees [][]string, current, idx int) int {
	result := 0
	for i := current + 1; i < len(trees); i++ {
		result++
		t, _ := strconv.Atoi(trees[current][idx])
		n, _ := strconv.Atoi(trees[i][idx])
		if n >= t {
			break
		}
	}
	return result
}

func countToRight(trees []string, start int) int {
	result := 0

	for i := start + 1; i < len(trees); i++ {
		result++
		n, _ := strconv.Atoi(trees[i])
		tt, _ := strconv.Atoi(trees[start])
		if n >= tt {
			break
		}
	}

	return result
}
