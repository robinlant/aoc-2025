package aoc

import (
	"bytes"
	"fmt"
	"strconv"
)

type Day5Solver struct{}

func (d *Day5Solver) GetDay() uint8 {
	return 5
}

func (d *Day5Solver) SolveOne(i []byte) (string, error) {
	ranges, ids, err := parseDay5input(i)
	if err != nil {
		return "", fmt.Errorf("failed parsind day5 input %s", err.Error())
	}

	count := 0
	for _, id := range ids {
		for _, r := range ranges {
			if r.Includes(id) {
				count++
				break
			}
		}
	}
	return strconv.Itoa(count), nil
}

func (d *Day5Solver) SolveTwo(i []byte) (string, error) {
	return "", nil
}

type indexRange struct {
	Start uint64
	Stop  uint64
}

func (i indexRange) Includes(n uint64) bool {
	return i.Start <= n && n <= i.Stop
}

func indexRangeFromBytes(s []byte) (indexRange, error) {
	if cnt := bytes.Count(s, []byte{'-'}); cnt != 1 {
		return indexRange{}, fmt.Errorf("error while creating index range '%s', expected 1 dash got %d", s, cnt)
	}
	parts := bytes.Split(s, []byte{'-'})
	start, err := strconv.ParseUint(string(parts[0]), 10, 64)
	if err != nil {
		return indexRange{}, err
	}
	stop, err := strconv.ParseUint(string(parts[1]), 10, 64)
	return indexRange{Start: start, Stop: stop}, err
}

func parseDay5input(i []byte) ([]indexRange, []uint64, error) {
	ranges := []indexRange{}
	ids := []uint64{}

	parsingIds := false

	for i, r := range enumerate(rows(i)) {
		if parsingIds {
			id, err := strconv.ParseUint(string(r), 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("error while parsing id '%s' at line %d - %s", r, i+1, err.Error())
			}
			ids = append(ids, id)
			continue
		}
		if len(r) == 0 {
			parsingIds = true
			continue
		}

		indRang, err := indexRangeFromBytes(r)
		if err != nil {
			return nil, nil, err
		}
		ranges = append(ranges, indRang)
	}
	return ranges, ids, nil
}
