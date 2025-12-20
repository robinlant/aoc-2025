package aoc

import (
	"bytes"
	"errors"
	"strconv"
)

type Day4Solver struct{}

func (d *Day4Solver) GetDay() uint8 {
	return 4
}

func (d *Day4Solver) SolveOne(i []byte) (string, error) {
	const empty = '.'
	const roll = '@'

	sum := 0
	return strconv.Itoa(sum), nil
}

func (d *Day4Solver) SolveTwo(i []byte) (string, error) {
	return "", nil
}

func getD4Matrix(i []byte) ([][]byte, error) {
	width := bytes.IndexAny(i, "\n")
	if width < 1 {
		return "", errors.New("unexpected problem format: new line as the first character")
	}
	matrix := [][]byte{}
	for r := range rows(i) {

	}
}
