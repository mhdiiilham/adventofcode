package cathoderaytube

import (
	"strconv"
	"strings"

	"github.com/mhdiiilham/adventofcode/pkg/input"
)

type Client struct {
	inputSource  string
	instructions []Instruction
	cycles       map[int]int
	message      string
	pixels       []bool
}

type Instruction struct {
	Type  string
	Value int
}

func New(inputSource string) (*Client, error) {
	c := &Client{inputSource: inputSource, cycles: map[int]int{}}

	rows, err := input.ReadByLines(c.inputSource, input.OneEnter)
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		command := strings.Split(row, " ")
		inst := Instruction{Type: command[0]}
		if len(command) > 1 {
			v, _ := strconv.Atoi(command[1])
			inst.Value = v
		}

		c.instructions = append(c.instructions, inst)
	}

	return c, nil
}

func (c *Client) GetSumOf6SignalStrengths() int {
	cycle := 0
	registerX := 1

	for _, instruction := range c.instructions {
		cycles := map[string]int{"noop": 1, "addx": 2}
		for i := 1; i <= cycles[instruction.Type]; i++ {
			if cycle%40 == 0 && cycle <= 220 {
				c.message = c.message + "\n"
			}

			if registerX-1 == cycle%40 || registerX == cycle%40 || registerX+1 == cycle%40 {
				c.message = c.message + "#"
			} else {
				c.message = c.message + "."
			}

			cycle++
			c.cycles[cycle] = registerX
		}
		if instruction.Type == "addx" {
			registerX += instruction.Value
		}
	}

	targetCycles := []int{20, 60, 100, 140, 180, 220}
	total := 0
	for _, cc := range targetCycles {
		sum := cc * c.cycles[cc]
		total += sum
	}

	return total
}

var alphabet = map[string]rune{
	".##.\n#..#\n#..#\n####\n#..#\n#..#": 'A',
	"###.\n#..#\n###.\n#..#\n#..#\n###.": 'B',
	".##.\n#..#\n#...\n#...\n#..#\n.##.": 'C',
	"####\n#...\n###.\n#...\n#...\n####": 'E',
	"####\n#...\n###.\n#...\n#...\n#...": 'F',
	".##.\n#..#\n#...\n#.##\n#..#\n.###": 'G',
	"#..#\n#..#\n####\n#..#\n#..#\n#..#": 'H',
	".###\n..#.\n..#.\n..#.\n..#.\n.###": 'I',
	"..##\n...#\n...#\n...#\n#..#\n.##.": 'J',
	"#..#\n#.#.\n##..\n#.#.\n#.#.\n#..#": 'K',
	"#...\n#...\n#...\n#...\n#...\n####": 'L',
	".##.\n#..#\n#..#\n#..#\n#..#\n.##.": 'O',
	"###.\n#..#\n#..#\n###.\n#...\n#...": 'P',
	"###.\n#..#\n#..#\n###.\n#.#.\n#..#": 'R',
	".###\n#...\n#...\n.##.\n...#\n###.": 'S',
	"#..#\n#..#\n#..#\n#..#\n#..#\n.##.": 'U',
	"#...\n#...\n.#.#\n..#.\n..#.\n..#.": 'Y',
	"####\n...#\n..#.\n.#..\n#...\n####": 'Z',
}

func (c *Client) GetTheEightCapitalLettersAppered() string {
	message := c.message
	lines := strings.Split(strings.TrimSpace(message), "\n")
	letters := []rune{}

	for i := 0; i < len(lines[0]); i += 5 {
		letterLines := []string{}
		for _, ln := range lines {
			letterLines = append(letterLines, ln[i:i+4])
		}
		letter := strings.Join(letterLines, "\n")
		letters = append(letters, alphabet[letter])
	}

	return string(letters)
}
