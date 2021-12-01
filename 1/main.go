package main

import (
	"bufio"
	"fmt"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
	"os"
)

const inputPath = "1/input.txt"

func CountMeasurementIncreases(measurements []int) int {
	var count int
	for i, measurement := range measurements {
		if i == 0 {
			continue
		}
		lastMeasurement := measurements[i-1]
		if measurement > lastMeasurement {
			count++
		}
	}
	return count
}

func Count3MeasurementWindowIncreases(measurements []int) int {
	var count int
	for i := 3; i < len(measurements); i++ {
		if measurements[i] > measurements[i-3] {
			count++
		}
	}
	return count
}

func main() {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	input, err := pkg.ReadIntsByDelimiter(bufio.NewReader(file), "\n")
	if err != nil {
		panic(err)
	}

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", CountMeasurementIncreases(input)) },
		func() string { return fmt.Sprintf("%v", Count3MeasurementWindowIncreases(input)) },
	)
}
