package aoc

import (
	"fmt"
	"strconv"
)

const empty = '.'
const roll = '@'

type Day4Solver struct{}

func (d *Day4Solver) GetDay() uint8 {
	return 4
}

func (d *Day4Solver) SolveOne(i []byte) (string, error) {
	matrix := getD4Matrix(i)

	sum := 0
	for rIndex, row := range matrix {
		for cIndex, char := range row {
			if char == '@' &&
				countAdjacentRolls(matrix, rIndex, cIndex) < 4 {
				sum++
				fmt.Printf("row: %d, col: %d\n", rIndex+1, cIndex+1)
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func (d *Day4Solver) SolveTwo(i []byte) (string, error) {
	return "", nil
}

func getD4Matrix(i []byte) [][]byte {
	matrix := [][]byte{}
	for r := range rows(i) {
		matrix = append(matrix, r)
	}
	return matrix
}

func countAdjacentRolls(m [][]byte, r, c int) uint8 {
	var count uint8 = 0

	for rIndex := r - 1; rIndex <= r+1; rIndex++ {
		if rIndex < 0 || rIndex > len(m)-1 {
			continue
		}
		for cIndex := c - 1; cIndex <= c+1; cIndex++ {
			if cIndex < 0 ||
				cIndex > len(m[0])-1 ||
				(cIndex == c && rIndex == r) {
				continue
			}
			char := m[rIndex][cIndex]
			if char == roll {
				count++
			}
		}
	}

	return count
}
