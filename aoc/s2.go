package aoc

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Day2Solver struct{}

type idRange struct {
	Start uint64
	Stop  uint64
}

func (d *Day2Solver) GetDay() uint8 {
	return 2
}

func (d *Day2Solver) SolveOne(i []byte) (string, error) {
	ranges, err := getRanges(i)
	if err != nil {
		return "", err
	}
	var sum uint64
	iterateIdRanges(ranges, func(n uint64) {
		id := strings.TrimLeft(strconv.FormatUint(n, 10), "0")
		if len(id) < 2 || len(id)%2 == 1 {
			return
		}
		if id[:len(id)/2] == id[len(id)/2:] {
			sum += n
		}
	})
	return strconv.FormatUint(sum, 10), nil
}

func (d *Day2Solver) SolveTwo(i []byte) (string, error) {
	ranges, err := getRanges(i)
	if err != nil {
		return "", err
	}
	var sum uint64
	var f func(id string, sub string) bool
	f = func(id string, sub string) bool {
		if len(sub) == 0 {
			return false
		}
		if id == strings.Repeat(sub, len(id)/len(sub)) {
			return true
		}
		return f(id, sub[:len(sub)-1])
	}
	iterateIdRanges(ranges, func(n uint64) {
		id := strconv.FormatUint(n, 10)
		if f(id, id[:len(id)/2]) {
			sum += n
		}
	})

	return strconv.FormatUint(sum, 10), nil
}

func iterateIdRanges(ranges []idRange, fn func(n uint64)) {
	for _, r := range ranges {
		for i := r.Start; i <= r.Stop; i++ {
			fn(i)
		}
	}
}

func getRanges(i []byte) ([]idRange, error) {
	trimI := bytes.Trim(i, " \n")
	rangeStrs := bytes.Split(trimI, []byte{','})
	ranges := make([]idRange, 0, len(rangeStrs))
	for _, rstr := range rangeStrs {
		rbSlice := bytes.Split(rstr, []byte{'-'})
		if len(rbSlice) != 2 {
			return nil, fmt.Errorf("expect len of 2 but got %d in '%v'", len(rbSlice), rbSlice)
		}
		start, err := strconv.ParseUint(string(rbSlice[0]), 10, 64)
		if err != nil {
			return nil, err
		}
		stop, err := strconv.ParseUint(string(rbSlice[1]), 10, 64)
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, idRange{Start: start, Stop: stop})
	}

	return ranges, nil
}
