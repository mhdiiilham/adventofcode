package suplystack

import (
	"strconv"
	"strings"

	"github.com/mhdiiilham/adventofcode/pkg/input"
)

type Instruction struct {
	From  int
	To    int
	Boxes int
}

type SuplyStack struct {
	InputSource  string
	StackOne     map[int][]string
	StackTwo     map[int][]string
	Instructions []Instruction
}

func New(inputSource string) *SuplyStack {
	return &SuplyStack{InputSource: inputSource, StackOne: map[int][]string{}, StackTwo: map[int][]string{}}
}

func (ss *SuplyStack) LoadInput() (err error) {
	rows, err := input.ReadByLines(ss.InputSource, input.TwoEnter)
	if err != nil {
		return err
	}

	instructions := strings.Split(rows[1], "\n")
	for _, instruction := range instructions {
		numbers := []int{}
		rows := strings.Split(instruction, " ")
		for _, row := range rows {
			n, err := strconv.Atoi(row)
			if err == nil {
				numbers = append(numbers, n)
			}
		}

		ss.Instructions = append(ss.Instructions, Instruction{
			Boxes: numbers[0],
			From:  numbers[1],
			To:    numbers[2],
		})
	}

	store := map[int]string{}
	stacks := rows[0]

	configurations := strings.Split(stacks, "\n")
	temp := []string{}
	for _, c := range configurations {
		rows := strings.Split(c, "")
		count := 1
		char := ""

		for _, row := range rows {
			if count == 4 {
				count = 1
				continue
			}

			if count == 1 && row != "[" {
				char += "["
			} else if count == 3 && row != "]" {
				char += "]"
			} else {
				char += row
			}
			count++
		}
		char = strings.ReplaceAll(char, "[", "")
		char = strings.ReplaceAll(char, "]", "")
		temp = append(temp, char)
	}

	stackRaw := temp[:len(temp)-1]
	for _, s := range stackRaw {
		list := strings.Split(s, "")
		for i, l := range list {
			store[i+1] += l
		}
	}

	for k, v := range store {
		store[k] = strings.ReplaceAll(v, " ", "")
		ss.StackOne[k] = strings.Split(store[k], "")
		ss.StackTwo[k] = strings.Split(store[k], "")
	}

	return nil
}

func (ss *SuplyStack) CrateMover9000() string {
	for _, instruction := range ss.Instructions {
		for i := 0; i < instruction.Boxes; i++ {
			group := ss.StackOne[instruction.From]
			topBox := group[0]
			ss.StackOne[instruction.From] = group[1:]
			insert := []string{topBox}
			ss.StackOne[instruction.To] = append(insert, ss.StackOne[instruction.To]...)
		}
	}

	result := ""
	for i := 0; i < len(ss.StackOne); i++ {
		result += ss.StackOne[i+1][0]
	}
	return result
}

func (ss *SuplyStack) CrateMover9001() string {
	var result string

	for _, instruction := range ss.Instructions {
		group := ss.StackTwo[instruction.From]
		boxToMoved := []string{}
		boxStayed := []string{}

		for _, v := range group {
			if len(boxToMoved) != instruction.Boxes {
				boxToMoved = append(boxToMoved, v)
			} else {
				boxStayed = append(boxStayed, v)
			}
		}
		ss.StackTwo[instruction.To] = append(boxToMoved, ss.StackTwo[instruction.To]...)
		ss.StackTwo[instruction.From] = boxStayed
	}

	for i := 0; i < len(ss.StackTwo); i++ {
		result += ss.StackTwo[i+1][0]
	}
	return result
}
