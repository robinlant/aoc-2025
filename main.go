package main

import (
	"flag"
	"fmt"
	"os"
	"robinlant/aoc-2025/aoc"
)

type args struct {
	Day   uint8
	Input string
}

func exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}

func parseArgs() *args {
	var input string
	var day uint64

	flag.StringVar(&input, "input", "", "input file of a problem")
	flag.Uint64Var(&day, "day", 0, "day of aoc-2025 puzzle")

	flag.Parse()

	return &args{
		Day:   uint8(day),
		Input: input,
	}
}

func main() {
	args := parseArgs()

	if args.Day == 0 || args.Input == "" {
		exit("please provide day via -day=<day> and input file via -input=<path>", 1)
	}

	i, err := os.ReadFile(args.Input)
	if err != nil {
		exit(fmt.Sprintf("error reading file '%s': %s", args.Input, err.Error()), 1)
	}

	a := aoc.NewAocService().
		AddHandler(&aoc.Day1Solver{})

	s1, s2, err := a.Solve(args.Day, i)

	if err != nil {
		exit(err.Error(), 1)
	}

	fmt.Printf("Day %d\n\n", args.Day)
	fmt.Printf("Problem 1: %s\nProblem 2: %s\n", s1, s2)
}
