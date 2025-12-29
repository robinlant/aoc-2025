package aoc

import (
	"bytes"
	"errors"
	"fmt"
	"math"
)

type Day8Solver struct{}

func (d *Day8Solver) GetDay() uint8 {
	return 8
}

func (d *Day8Solver) SolveOne(i []byte) (string, error) {
	const pairs = 1000
	points, err := parseDay8input(i)
	if err != nil {
		return "", err
	}

	unused(points, err)
	return "", nil
}

func (d *Day8Solver) SolveTwo(i []byte) (string, error) {
	return "", nil
}

func parseDay8input(i []byte) ([]point, error) {
	points := []point{}
	for i, r := range enumerate(rows(i)) {
		cords := bytes.Split(r, []byte(","))
		if len(cords) != 3 {
			return nil, fmt.Errorf("line %d expect row '%s' to have 3 cordinates to be separated by commas", i, r)
		}
		convertedCords, err := parseBytesIntoInt(cords)
		if err != nil {
			return nil, fmt.Errorf("error parsing day 8 input: %s", err.Error())
		}
		points = append(points, point{
			X: convertedCords[0],
			Y: convertedCords[1],
			Z: convertedCords[2],
		})
	}
	return points, nil
}

type pointLike interface {
	Distance(point) float64
	PointArr() []point
}

type point struct {
	X int
	Y int
	Z int
}

func (p point) PointArr() []point {
	return []point{p}
}

func (p1 point) Distance(p2 point) float64 {
	squareDif := func(x int, y int) int {
		return (x - y) ^ 2
	}
	x := squareDif(p1.X, p2.X)
	y := squareDif(p1.Y, p2.Y)
	z := squareDif(p1.Z, p2.Z)

	return math.Sqrt(float64(x + y + z))
}

type pointConnection struct {
	Points []point
}

func (c pointConnection) Distance(p point) float64 {
	if len(c.Points) == 0 {
		panic(errors.New("cannot count distance as pointConnection has no points"))
	}
	min := c.Points[0].Distance(p)
	for i := 1; i < len(c.Points); i++ {
		d := c.Points[i].Distance(p)
		if d < min {
			min = d
		}
	}
	return min
}

func (c pointConnection) PointArr() []point {
	return c.Points
}

func (c *pointConnection) Add(p pointLike) {
	c.Points = append(c.Points, p.PointArr()...)
}

func newConnection(pl ...pointLike) pointConnection {
	con := pointConnection{}
	for _, p := range pl {
		con.Add(p)
	}
	return con
}
