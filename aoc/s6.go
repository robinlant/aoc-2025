package aoc

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

const operationD6Row = 4

type Day6Solver struct{}

func (d *Day6Solver) GetDay() uint8 {
	return 6
}

func (d *Day6Solver) SolveOne(i []byte) (string, error) {
	problems, err := parseDay6problem(i)
	if err != nil {
		return "", fmt.Errorf("error while parsing day 6 %s", err)
	}
	sum, err := solveProblems(problems)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(sum, 10), nil
}

// TODO finish
func (d *Day6Solver) SolveTwo(i []byte) (string, error) {
	problems, err := parseDay6problem(i)
	if err != nil {
		return "", fmt.Errorf("error while parsing day 6 %s", err)
	}
	for i, p := range problems {
		problems[i] = adaptToCephalopod(p)
	}
	sum, err := solveProblems(problems)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(sum, 10), nil
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

func adaptToCephalopod(p mathProblem) mathProblem {
	// TODO implement this function as a sort of adapter taht will adapt
	//	celapthore math to normal math problems
	// numStrins := make([]byte, 0, len(p.Numbers))
	// _ := numStrings
	return p
}

func parseDay6problem(i []byte) ([]mathProblem, error) {
	i = regexp.MustCompile(`[ \t]+`).ReplaceAll(i, []byte{' '})
	rows := rowsSlice(i)
	problems := []mathProblem{}
	for o := range bytes.SplitSeq(rows[operationD6Row], []byte{' '}) {
		if len(o) != 1 {
			return nil, fmt.Errorf("expect operation to be only one byte but got %d (%s)", len(o), o)
		}
		problems = append(problems, mathProblem{Operation: o[0]})
	}
	for i := range operationD6Row {
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
