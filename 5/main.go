package main

import (
	"bufio"
	"fmt"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
	"log"
	"os"
	"strings"

	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
)

const inputPath = "5/input.txt"

func getPointsInYAxis(p1, p2 types.Point) []types.Point {
	var points []types.Point

	// validate
	if p1.X != p2.X {
		return nil
	}

	yDirection := 1
	if p1.Y > p2.Y {
		yDirection = -1
	}

	for y := p1.Y; y != p2.Y; y += yDirection {
		points = append(points, types.Point{X: p1.X, Y: y})
	}

	return append(points, p2)
}

func getPointsInXAxis(p1, p2 types.Point) []types.Point {
	var points []types.Point

	// validate
	if p1.Y != p2.Y {
		return nil
	}

	xDirection := 1
	if p1.X > p2.X {
		xDirection = -1
	}

	for x := p1.X; x != p2.X; x += xDirection {
		points = append(points, types.Point{X: x, Y: p1.Y})
	}

	return append(points, p2)
}

func getPointsInDiagonal(p1, p2 types.Point) []types.Point {
	var points []types.Point

	xDirection := 1
	if p1.X > p2.X {
		xDirection = -1
	}

	yDirection := 1
	if p1.Y > p2.Y {
		yDirection = -1
	}

	// validate
	if xDirection*(p1.X-p2.X) != yDirection*(p1.Y-p2.Y) {
		return nil
	}

	x := p1.X
	y := p1.Y
	for y != p2.Y && x != p2.X {
		points = append(points, types.Point{X: x, Y: y})
		x += xDirection
		y += yDirection
	}

	return append(points, p2)
}

func CheckNumberOfOverlaps(lines []types.Line, withDiagonals bool) int {
	var overlaps int

	mapPointsToAppears := map[string]int{}
	for _, l := range lines {
		var pointsInBetween = getPointsInXAxis(l.Start, l.End)
		pointsInBetween = append(pointsInBetween, getPointsInYAxis(l.Start, l.End)...)
		if withDiagonals {
			pointsInBetween = append(pointsInBetween, getPointsInDiagonal(l.Start, l.End)...)
		}

		for _, p := range pointsInBetween {
			_, exists := mapPointsToAppears[p.ToString()]
			if !exists {
				mapPointsToAppears[p.ToString()] = 1
			} else {
				mapPointsToAppears[p.ToString()]++
				if mapPointsToAppears[p.ToString()] == 2 {
					overlaps++
				}
			}
		}
	}

	return overlaps
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

	var lines []types.Line
	for _, line := range fileLines {
		splitter := strings.Split(line, " -> ")
		p1, err := types.PointFromString(splitter[0])
		pkg.PanicErr(err)
		p2, err := types.PointFromString(splitter[1])
		pkg.PanicErr(err)

		lines = append(lines, types.Line{Start: p1, End: p2})
	}

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", CheckNumberOfOverlaps(lines, false)) },
		func() string { return fmt.Sprintf("%v", CheckNumberOfOverlaps(lines, true)) },
	)
}
