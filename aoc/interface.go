package aoc

import "fmt"

type ProblemHandler interface {
	GetDay() uint8
	SolveOne(i []byte) (string, error)
	SolveTwo(i []byte) (string, error)
}

type AocService struct {
	handlers map[uint8]ProblemHandler
}

func NewAocService() *AocService {
	m := make(map[uint8]ProblemHandler, 24)

	return &AocService{handlers: m}
}

func (a *AocService) Solve(day uint8, input []byte) (string, string, error) {
	h, ok := a.handlers[day]
	if !ok {
		return "", "", fmt.Errorf("handler for day %d doesn't exist", day)
	}
	s1, err := h.SolveOne(input)
	if err != nil {
		return "", "", fmt.Errorf("got error while solving day %d problem one: %s", day, err.Error())
	}
	s2, err := h.SolveTwo(input)
	if err != nil {
		return "", "", fmt.Errorf("got error while solving day %d problem two: %s", day, err.Error())
	}
	return s1, s2, nil
}

// adds a day handler to AocService handler. Adding a day, that already exist, overwrites it
func (a *AocService) AddHandler(h ProblemHandler) *AocService {
	a.handlers[h.GetDay()] = h
	return a
}
