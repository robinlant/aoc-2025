package aoc

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Day1Solver struct{}

type instruction struct {
	Direction byte
	Count     int
}

func getInstructions(i []byte) ([]instruction, error) {
	re := regexp.MustCompile(`[LR]\d+`)

	m := re.FindAll(i, -1)
	if m == nil {
		return []instruction{}, errors.New("no instructions were found")
	}
	instr := make([]instruction, 0, len(m))
	for _, v := range m {
		cnt, err := strconv.ParseInt(string(v[1:]), 10, 32)
		if err != nil {
			return []instruction{}, err
		}
		instr = append(instr, instruction{Count: int(cnt), Direction: v[0]})
	}

	return instr, nil
}

func (d *Day1Solver) GetDay() uint8 {
	return 1
}

func (d *Day1Solver) SolveOne(i []byte) (string, error) {
	inst, err := getInstructions(i)
	if err != nil {
		return "", err
	}

	resultCount := 0
	count := 50
	for _, v := range inst {
		switch v.Direction {
		case 'R':
			count += v.Count
		case 'L':
			count -= v.Count
		default:
			return "", fmt.Errorf("unknown direction %c", v.Direction)
		}
		if count%100 == 0 {
			resultCount += 1
		}
	}

	return strconv.Itoa(resultCount), nil
}

func (d *Day1Solver) SolveTwo(i []byte) (string, error) {
	inst, err := getInstructions(i)
	if err != nil {
		return "", err
	}

	resultCount := 0
	currentPos := 50

	for _, v := range inst {
		for range v.Count {
			if v.Direction == 'R' {
				currentPos++
			} else if v.Direction == 'L' {
				currentPos--
			} else {
				return "", fmt.Errorf("unknown direction %c", v.Direction)
			}

			if currentPos%100 == 0 {
				resultCount++
			}
		}
	}

	return strconv.Itoa(resultCount), nil
}
