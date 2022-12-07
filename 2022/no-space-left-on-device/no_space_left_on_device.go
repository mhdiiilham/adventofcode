package nospaceleftondevice

import (
	"path"
	"strconv"
	"strings"

	"github.com/mhdiiilham/adventofcode/pkg/input"
)

type Client struct {
	inputSource string
	rows        []string
	fileSystem  map[string]int
}

func New(inputSource string) (c *Client, err error) {
	rows, err := input.ReadByLines(inputSource, input.OneEnter)
	if err != nil {
		return nil, err
	}

	client := Client{inputSource: inputSource, rows: rows}
	client.fileSystem = client.crawlFileSystem()
	return &client, nil

}

func (c *Client) crawlFileSystem() map[string]int {
	pwd := ""
	directories := map[string]int{}
	for _, row := range c.rows {
		if strings.HasPrefix(row, "$ cd") {
			pwd = path.Join(pwd, strings.ReplaceAll(row, "$ cd ", ""))
		}

		if !strings.HasPrefix(row, "$") && !strings.HasPrefix(row, "dir") {
			value := strings.Split(row, " ")
			fileName := value[1]
			size, _ := strconv.Atoi(value[0])
			directories[path.Join(pwd, fileName)] = size
		}

	}
	return directories
}

func (c *Client) GetTotalSizeOfDirectories() int {
	sizeOfDirectory := getSizeOfDirectory(c.fileSystem)
	sum := 0
	for _, v := range sizeOfDirectory {
		if v > 0 && v < 100000 {
			sum += v
		}
	}
	return sum
}

func (c *Client) GetTotalSizeDirectoryNeedToDelete() int {
	sizeOfDirectory := getSizeOfDirectory(c.fileSystem)
	targetFree := 30000000 - (70000000 - sizeOfDirectory["/"])
	currentMin := 70000000

	for _, v := range sizeOfDirectory {
		if v > targetFree && v < currentMin {
			currentMin = v
		}
	}
	return currentMin
}

func getSizeOfDirectory(fileSystem map[string]int) map[string]int {
	dirs := map[string]int{}

	for fpath, size := range fileSystem {
		for dir := path.Dir(fpath); ; dir = path.Dir(dir) {
			dirs[dir] += size
			if dir == "/" {
				break
			}
		}
	}

	return dirs
}
