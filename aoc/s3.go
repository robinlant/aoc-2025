package aoc

import (
	"strconv"
)

type Day3Solver struct{}

func (d *Day3Solver) GetDay() uint8 {
	return 3
}

func (d *Day3Solver) SolveOne(i []byte) (string, error) {
	sum := 0
	for row := range rows(i) {
		first := 0
		second := 0
		for i, r := range row {
			j, err := digiByteToInt(r)
			if err != nil {
				return "", err
			}
			if j > first && i < len(row)-1 {
				first = j
				second = 0
				continue
			}
			if j > second {
				second = j
			}
		}
		sum += first*10 + second
	}

	return strconv.Itoa(sum), nil
}

func (d *Day3Solver) SolveTwo(i []byte) (string, error) {
	const digits int = 2
	return maxJoltage(i, digits)
}

// TODO debug wiht digits 2 and wait till result is the same as problem 1
func maxJoltage(i []byte, digits int) (string, error) {
	var sum uint64
	for row := range rows(i) {
		joltageSlice := make([]uint64, digits)
		for index, char := range row {
			v, err := digiByteToInt(char)
			if err != nil {
				return "", err
			}
			joltage := uint64(v)

			var swapped bool
			for i, v := range joltageSlice {
				if index+digits-i >= len(row) {
					continue
				}
				if swapped {
					joltageSlice[i] = 0
					continue
				}
				if joltage > v {
					joltageSlice[i] = joltage
					swapped = true
				}
			}
		}
		for i, j := range joltageSlice {
			sum += j*10 ^ uint64(digits-i-1)
		}
	}

	return strconv.FormatUint(sum, 10), nil
}
