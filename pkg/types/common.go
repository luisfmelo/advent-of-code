package types

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

type MatrixElement struct {
	LineIndex, ColumnIndex int
}

func (m MatrixElement) ToString() string {
	return fmt.Sprintf("(%v,%v)", m.LineIndex, m.ColumnIndex)
}

type Point struct {
	X, Y int
}

func (p Point) ToString() string {
	return fmt.Sprintf("%v,%v", p.X, p.Y)
}

// PointFromString transforms a string in form of "X,Y" into a Point.
func PointFromString(str string) (Point, error) {
	splitter := strings.Split(str, ",")

	x, err := strconv.Atoi(splitter[0])
	if err != nil {
		return Point{}, err
	}

	y, err := strconv.Atoi(splitter[1])
	if err != nil {
		return Point{}, err
	}

	return Point{X: x, Y: y}, nil
}

type Line struct {
	Start, End Point
}
