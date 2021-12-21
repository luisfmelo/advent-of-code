package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputPath = "13/input.txt"

type foldDirection int

const (
	dot = 1

	foldVertical   foldDirection = 1
	foldHorizontal foldDirection = 2
)

type Instruction struct {
	fType  foldDirection
	cutoff int
}

func NewInstruction(s string) Instruction {
	var err error
	var i Instruction
	if strings.Contains(s, "fold along y=") {
		i.fType = foldHorizontal
		i.cutoff, err = strconv.Atoi(strings.Replace(s, "fold along y=", "", -1))
		pkg.PanicErr(err)
	} else if strings.Contains(s, "fold along x") {
		i.fType = foldVertical
		i.cutoff, err = strconv.Atoi(strings.Replace(s, "fold along x=", "", -1))
		pkg.PanicErr(err)
	} else {
		pkg.PanicErr(errors.New("fold type unknown"))
	}
	return i
}

func newMatrix(initialDots []types.Point) types.Matrix {
	var maxLineIdx, maxColumnsIdx int
	for _, p := range initialDots {
		if p.X > maxColumnsIdx {
			maxColumnsIdx = p.X
		}
		if p.Y > maxLineIdx {
			maxLineIdx = p.Y
		}
	}

	m := types.NewMatrix(maxLineIdx+1, maxColumnsIdx+1)
	for _, p := range initialDots {
		m[p.Y][p.X] = dot
	}
	return m
}

func countDots(m types.Matrix) int {
	var c int
	for _, line := range m {
		for _, p := range line {
			if p == dot {
				c++
			}
		}
	}
	return c
}

func fold(m types.Matrix, instruction Instruction) types.Matrix {
	var newM types.Matrix

	switch instruction.fType {
	case foldVertical:
		var (
			rightIdx = instruction.cutoff + 1
			leftIdx  = instruction.cutoff - 1
		)

		newM = types.NewMatrix(m.NumberOfRows(), m.NumberOfColumns()/2)
		for rightIdx < m.NumberOfColumns() && leftIdx >= 0 {
			for rowIdx := 0; rowIdx < newM.NumberOfRows(); rowIdx++ {
				if m[rowIdx][rightIdx] == dot || m[rowIdx][leftIdx] == dot {
					newM[rowIdx][leftIdx] = dot
				}
			}

			rightIdx++
			leftIdx--
		}

	case foldHorizontal:
		var (
			downIdx = instruction.cutoff + 1
			upIdx   = instruction.cutoff - 1
		)

		newM = types.NewMatrix(m.NumberOfRows()/2, m.NumberOfColumns())
		for downIdx < m.NumberOfRows() && upIdx >= 0 {
			for colIdx := 0; colIdx < newM.NumberOfColumns(); colIdx++ {
				if m[downIdx][colIdx] == dot || m[upIdx][colIdx] == dot {
					newM[upIdx][colIdx] = dot
				}
			}

			downIdx++
			upIdx--
		}
	}

	return newM
}

func CountNumberOfDots(initialDots []types.Point, printCode bool, instructions ...Instruction) int {
	m := newMatrix(initialDots)

	for _, instruction := range instructions {
		m = fold(m, instruction)
	}

	if printCode {
		m.PrintWithMapping(map[int]rune{0: ' ', 1: '#'})
	}

	return countDots(m)
}

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error occurred: %v", err)
		}
	}()

	file, err := os.Open(inputPath)
	pkg.PanicErr(err)

	r := bufio.NewReader(file)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	fileLines, err := pkg.ReadLines(scanner)
	pkg.PanicErr(err)

	var instructions []Instruction
	var initialDots []types.Point
	initialDotsTerminated := false
	for _, line := range fileLines {
		if line == "" {
			initialDotsTerminated = true
			continue
		}

		if initialDotsTerminated {
			instructions = append(instructions, NewInstruction(line))
		} else {
			p, err := types.PointFromString(line)
			pkg.PanicErr(err)
			initialDots = append(initialDots, p)
		}
	}

	pkg.RunWithTime(
		func() string {
			return fmt.Sprintf("%v", CountNumberOfDots(initialDots, false, instructions[0]))
		},
		func() string {
			return fmt.Sprintf("%v", CountNumberOfDots(initialDots, true, instructions...))
		},
	)
}
