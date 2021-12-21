package types

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func NewMatrix(nLines, nColumns int) Matrix {
	m := Matrix{}
	for i := 0; i < nLines; i++ {
		m = append(m, make([]int, nColumns))
	}

	return m
}

func (m Matrix) NumberOfRows() int {
	return len(m)
}

func (m Matrix) NumberOfColumns() int {
	if m.NumberOfRows() == 0 {
		return 0
	}

	return len(m[0])
}

func (m Matrix) PrintWithMapping(mapping map[int]rune) {
	for i := range m {
		for j := range m[i] {
			r, exists := mapping[m[i][j]]
			if !exists {
				r = '?'
			}
			fmt.Printf("%s", string(r))
		}
		fmt.Println("")
	}
}

// GetAdjacentOfElement will return all elements adjacent to the center element.
// It will include diagonal elements
func (m Matrix) GetAdjacentOfElement(c MatrixElement) []MatrixElement {
	var adjacentElements []MatrixElement
	for i := c.LineIndex - 1; i <= c.LineIndex+1; i++ {
		if i < 0 || i > len(m)-1 {
			continue
		}
		for j := c.ColumnIndex - 1; j <= c.ColumnIndex+1; j++ {
			if j < 0 || j > len(m[i])-1 || i == c.LineIndex && j == c.ColumnIndex {
				continue
			}
			adjacentElements = append(adjacentElements, MatrixElement{LineIndex: i, ColumnIndex: j})
		}
	}

	return adjacentElements
}

type MatrixElement struct {
	LineIndex, ColumnIndex int
}

func (e MatrixElement) ToString() string {
	return fmt.Sprintf("(%v,%v)", e.LineIndex, e.ColumnIndex)
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
