package aoc

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

type Day6Solver struct{}

func (d *Day6Solver) GetDay() uint8 {
	return 6
}

func (d *Day6Solver) SolveOne(i []byte) (string, error) {
	problems, err := parseDay6problem1(i)
	if err != nil {
		return "", fmt.Errorf("error while parsing day 6 %s", err)
	}
	sum, err := solveProblems(problems)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(sum, 10), nil
}

func (d *Day6Solver) SolveTwo(i []byte) (string, error) {
	return "", nil
}

func solveProblems(p []mathProblem) (int64, error) {
	var sum int64
	for _, p := range p {
		switch p.Operation {
		case '+':
			for _, n := range p.Numbers {
				sum += int64(n)
			}
		case '*':
			mult := 1
			for _, n := range p.Numbers {
				mult *= n
			}
			sum += int64(mult)
		default:
			return 0, fmt.Errorf("got unexpected operation: %q", p.Operation)
		}
	}
	return sum, nil
}

func parseDay6problem1(i []byte) ([]mathProblem, error) {
	const operationRow = 4
	i = regexp.MustCompile(`[ \t]+`).ReplaceAll(i, []byte{' '})
	rows := rowsSlice(i)
	problems := []mathProblem{}
	for o := range bytes.SplitSeq(rows[operationRow], []byte{' '}) {
		if len(o) != 1 {
			return nil, fmt.Errorf("expect operation to be only one byte but got %d (%s)", len(o), o)
		}
		problems = append(problems, mathProblem{Operation: o[0]})
	}
	for i := range operationRow {
		splitRow := bytes.Split(bytes.TrimSpace(rows[i]), []byte{' '})
		for i, num := range splitRow {
			num, err := strconv.ParseInt(string(num), 10, 32)
			if err != nil {
				return nil, err
			}
			problems[i].Numbers = append(problems[i].Numbers, int(num))
		}
	}

	return problems, nil
}

type mathProblem struct {
	Numbers   []int
	Operation byte
}
