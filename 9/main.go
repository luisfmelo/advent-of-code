package main

import (
	"bufio"
	"fmt"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
	"log"
	"os"
	"sort"
)

const inputPath = "9/input.txt"

func GetNeighbours(lavaTunnelHeights types.Matrix, c types.MatrixElement) []types.MatrixElement {
	var neighbours []types.MatrixElement

	if c.LineIndex > 0 && lavaTunnelHeights[c.LineIndex-1][c.ColumnIndex] > -1 {
		neighbours = append(neighbours, types.MatrixElement{LineIndex: c.LineIndex - 1, ColumnIndex: c.ColumnIndex})
	}
	if c.LineIndex < len(lavaTunnelHeights)-1 && lavaTunnelHeights[c.LineIndex+1][c.ColumnIndex] > -1 {
		neighbours = append(neighbours, types.MatrixElement{LineIndex: c.LineIndex + 1, ColumnIndex: c.ColumnIndex})
	}
	if c.ColumnIndex > 0 && lavaTunnelHeights[c.LineIndex][c.ColumnIndex-1] > -1 {
		neighbours = append(neighbours, types.MatrixElement{LineIndex: c.LineIndex, ColumnIndex: c.ColumnIndex - 1})
	}
	if c.ColumnIndex < len(lavaTunnelHeights[c.LineIndex])-1 && lavaTunnelHeights[c.LineIndex][c.ColumnIndex+1] > -1 {
		neighbours = append(neighbours, types.MatrixElement{LineIndex: c.LineIndex, ColumnIndex: c.ColumnIndex + 1})
	}

	return neighbours
}

func isLavaTunnelLowLevel(lavaTunnelHeights types.Matrix, c types.MatrixElement) bool {
	cHeight := lavaTunnelHeights[c.LineIndex][c.ColumnIndex]

	if cHeight == 9 {
		return false
	}

	// up
	if c.LineIndex > 0 && lavaTunnelHeights[c.LineIndex-1][c.ColumnIndex] <= cHeight && lavaTunnelHeights[c.LineIndex-1][c.ColumnIndex] > -1 {
		return false
	}
	// down
	if c.LineIndex < len(lavaTunnelHeights)-1 && lavaTunnelHeights[c.LineIndex+1][c.ColumnIndex] <= cHeight && lavaTunnelHeights[c.LineIndex+1][c.ColumnIndex] > -1 {
		return false
	}
	// right
	if c.ColumnIndex > 0 && lavaTunnelHeights[c.LineIndex][c.ColumnIndex-1] <= cHeight && lavaTunnelHeights[c.LineIndex][c.ColumnIndex-1] > -1 {
		return false
	}
	// left
	if c.ColumnIndex < len(lavaTunnelHeights[c.LineIndex])-1 && lavaTunnelHeights[c.LineIndex][c.ColumnIndex+1] <= cHeight && lavaTunnelHeights[c.LineIndex][c.ColumnIndex+1] > -1 {
		return false
	}

	return true
}

func getLavaTunnelLowLevels(lavaTunnelHeights types.Matrix) []types.MatrixElement {
	var lowLevels []types.MatrixElement
	for lineIndex, line := range lavaTunnelHeights {
		for columnIndex := range line {
			center := types.MatrixElement{LineIndex: lineIndex, ColumnIndex: columnIndex}
			if isLavaTunnelLowLevel(lavaTunnelHeights, center) {
				lowLevels = append(lowLevels, types.MatrixElement{LineIndex: lineIndex, ColumnIndex: columnIndex})
			}
		}
	}

	return lowLevels
}

func GetBasinSize(lavaTunnelHeights types.Matrix, c types.MatrixElement) int {
	queue := types.NewMatrixElementQueue()
	queue.Push(c)

	var basinSize int
	for queue.Size() > 0 {
		c = queue.Pop()
		lavaTunnelHeights[c.LineIndex][c.ColumnIndex] = -1
		basinSize++
		for _, n := range GetNeighbours(lavaTunnelHeights, c) {
			if isLavaTunnelLowLevel(lavaTunnelHeights, n) {
				queue.Push(n)
			}
		}
	}

	return basinSize
}

func SumRiskLevelsOfLavaTunnel(lavaTunnelHeights types.Matrix) int {
	var sumRiskLevels int
	for _, lowLevelPoint := range getLavaTunnelLowLevels(lavaTunnelHeights) {
		sumRiskLevels += 1 + lavaTunnelHeights[lowLevelPoint.LineIndex][lowLevelPoint.ColumnIndex]
	}

	return sumRiskLevels
}

func Multiply3GreatestBasins(lavaTunnelHeights types.Matrix) int {
	var basins []int
	for _, lowLevelPoint := range getLavaTunnelLowLevels(lavaTunnelHeights) {
		basins = append(basins, GetBasinSize(pkg.CopyMatrix(lavaTunnelHeights), lowLevelPoint))
	}

	sort.Ints(basins)

	result := 1
	for _, biggestBasin := range basins[len(basins)-3:] {
		result *= biggestBasin
	}

	return result
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

	lavaTunnelHeights, err := pkg.ReadDigitMatrixWithoutSeparator(scanner)
	pkg.PanicErr(err)

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", SumRiskLevelsOfLavaTunnel(lavaTunnelHeights)) },
		func() string { return fmt.Sprintf("%v", Multiply3GreatestBasins(lavaTunnelHeights)) },
	)
}
