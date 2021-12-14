package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
)

const inputPath = "11/input.txt"

const flash = 100

func increaseEnergy(m types.Matrix, c types.MatrixElement) (types.Matrix, bool) {
	switch m[c.LineIndex][c.ColumnIndex] {
	case flash:
		return m, false
	case 9:
		m[c.LineIndex][c.ColumnIndex] = flash
	default:
		m[c.LineIndex][c.ColumnIndex]++
	}

	return m, m[c.LineIndex][c.ColumnIndex] == flash
}

func increaseEnergyOfAll(m types.Matrix) (types.Matrix, []types.MatrixElement) {
	var flashedOctapuses []types.MatrixElement
	var flashed bool
	for i, line := range m {
		for j := range line {
			octapus := types.MatrixElement{LineIndex: i, ColumnIndex: j}
			m, flashed = increaseEnergy(m, octapus)
			if flashed {
				flashedOctapuses = append(flashedOctapuses, octapus)
			}
		}
	}

	return m, flashedOctapuses
}

func countFlashes(m types.Matrix) int {
	var flashes int
	for _, line := range m {
		for _, energy := range line {
			if energy == flash {
				flashes++
			}
		}
	}

	return flashes
}

func rebootFlashes(m types.Matrix) types.Matrix {
	for i, line := range m {
		for j, energy := range line {
			if energy == flash {
				m[i][j] = 0
			}
		}
	}

	return m
}

func newStep(m types.Matrix) (types.Matrix, int) {
	var (
		flashes          int
		flashed          bool
		flashedOctapuses []types.MatrixElement
	)

	// Increase energy of all octapuses & Get all octapuses with energy level > 9
	m, flashedOctapuses = increaseEnergyOfAll(m)

	// Increase energy level of adjacent octapuses that flashed
	queue := types.NewMatrixElementQueue()
	queue.PushMultiple(flashedOctapuses)
	for queue.Size() > 0 {
		flashedOctapus := queue.Pop()
		for _, adjacentOctapus := range m.GetAdjacentOfElement(flashedOctapus) {
			m, flashed = increaseEnergy(m, adjacentOctapus)
			if flashed {
				queue.Push(adjacentOctapus)
			}
		}
	}

	// Get number of flashes in this step
	flashes = countFlashes(m)

	// Set all flashed octapuses energy to 0
	m = rebootFlashes(m)

	return m, flashes
}

func CountNumberOfFlashes(m types.Matrix, steps int) int {
	var flashes, totalFlashes int

	for step := 0; step < steps; step++ {
		m, flashes = newStep(m)
		totalFlashes += flashes
	}

	return totalFlashes
}

func CountNumberOfStepsUntilSimultaneousFlashes(m types.Matrix) int {
	var (
		flashes int
		step    int
	)

	for step = 0; flashes != 100; step++ {
		m, flashes = newStep(m)
	}

	return step
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

	matrix, err := pkg.ReadDigitMatrixWithoutSeparator(scanner)
	pkg.PanicErr(err)

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", CountNumberOfFlashes(matrix, 100)) },
		func() string { return fmt.Sprintf("%v", CountNumberOfStepsUntilSimultaneousFlashes(matrix)) },
	)
}
