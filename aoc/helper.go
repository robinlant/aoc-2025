package aoc

import (
	"fmt"
	"iter"
)

func rows(input []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		r := []byte{}
		for i := range input {
			if input[i] != '\n' {
				r = append(r, input[i])
				continue
			}
			if !yield(r) {
				return
			}
			r = []byte{}
		}
	}
}

func digiByteToInt(b byte) (int, error) {
	if b < '0' || b > '9' {
		return 0, fmt.Errorf("%c is not a digit", b)
	}
	return int(b - '0'), nil
}
