package aoc

import (
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
			if char == roll &&
				countAdjacentRolls(matrix, rIndex, cIndex) < 4 {
				sum++
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func (d *Day4Solver) SolveTwo(i []byte) (string, error) {
	matrix := getD4Matrix(i)
	type point struct {
		Row int
		Col int
	}

	sum := 0
	for true {
		rmPoints := []point{}

		for rIndex, row := range matrix {
			for cIndex, char := range row {
				if char == roll &&
					countAdjacentRolls(matrix, rIndex, cIndex) < 4 {
					rmPoints = append(rmPoints, point{Row: rIndex, Col: cIndex})
					sum++
				}
			}
		}

		if len(rmPoints) == 0 {
			break
		}
		for _, p := range rmPoints {
			matrix[p.Row][p.Col] = empty
		}
	}

	return strconv.Itoa(sum), nil
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
