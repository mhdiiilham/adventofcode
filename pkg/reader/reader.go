package reader

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type reader struct {
	FilePath  string
	rawResult []string
}

func NewReader(filepath string) *reader {
	return &reader{FilePath: filepath}
}

func (r *reader) ReadInput() (err error) {
	var b []byte

	b, err = ioutil.ReadFile(r.FilePath)
	if err != nil {
		log.Printf("error reading file: %v", err)
		return
	}

	r.rawResult = strings.Split(string(b), "\n")
	return nil
}

func (r *reader) GetResultAsSliceOfInt() (results []int, err error) {
	for _, in := range r.rawResult {
		var i int
		i, err = strconv.Atoi(in)
		if err != nil {
			return
		}

		results = append(results, i)
	}

	return
}
