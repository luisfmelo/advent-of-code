package main

import (
	"bufio"
	"fmt"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
	"log"
	"os"
	"sort"
)

const inputPath = "7/input.txt"

var fuelCostCache = map[int]int{}

func calculateFuelCostOnStep(steps int) int {
	if cost, exists := fuelCostCache[steps]; exists {
		return cost
	}
	cost := 0
	for s := steps; s > 0; s-- {
		cost += s
	}

	fuelCostCache[steps] = cost

	return cost
}

func CalculateLeastAmountOfFuelWhenBurnedAtConstantRate(crabPositions []int) int {
	sort.Ints(crabPositions)
	median := crabPositions[len(crabPositions)/2]

	var fuel int
	for _, hPosition := range crabPositions {
		fuel += pkg.IntAbs(hPosition - median)
	}

	return fuel
}

func CalculateLeastAmountOfFuelWhenBurnedAtIncreasingRate(crabPositions []int) int {
	sum := 0
	for _, hPosition := range crabPositions {
		sum += hPosition
	}

	avg := sum / len(crabPositions)
	mapFuel := map[int]int{avg - 1: 0, avg: 0, avg + 1: 0}
	for _, hPosition := range crabPositions {
		for targetPos := range mapFuel {
			mapFuel[targetPos] += calculateFuelCostOnStep(pkg.IntAbs(hPosition - targetPos))
		}
	}

	minFuel := mapFuel[avg]
	for _, fuel := range mapFuel {
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
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

	crabPositions, err := pkg.ReadIntegersInLine(scanner, ",")
	pkg.PanicErr(err)

	pkg.RunWithTime(
		func() string {
			return fmt.Sprintf("%v", CalculateLeastAmountOfFuelWhenBurnedAtConstantRate(crabPositions))
		},
		func() string {
			return fmt.Sprintf("%v", CalculateLeastAmountOfFuelWhenBurnedAtIncreasingRate(crabPositions))
		},
	)
}
