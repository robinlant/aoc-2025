package aoc

import (
	"fmt"
	"iter"
)

func rows(input []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		start := 0
		for i := 0; i < len(input); i++ {
			if input[i] == '\n' {
				if !yield(input[start:i]) {
					return
				}
				start = i + 1
			}
		}
		if start < len(input) {
			yield(input[start:])
		}
	}
}

func digiByteToInt(b byte) (int, error) {
	if b < '0' || b > '9' {
		return 0, fmt.Errorf("%c is not a digit", b)
	}
	return int(b - '0'), nil
}
