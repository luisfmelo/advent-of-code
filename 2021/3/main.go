package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
)

const inputPath = "2021/3/input.txt"

func mostCommonByteInPosition(diagnosticReport []string, pos int) rune {
	var ones, zeros int
	for _, l := range diagnosticReport {
		if l[pos] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	if ones >= zeros {
		return '1'
	}

	return '0'
}

func CalculatePowerConsumption(diagnosticReport []string) int64 {
	var gamaRateStr, epsilonRateStr string
	for index := range diagnosticReport[0] {
		mostCommonByte := mostCommonByteInPosition(diagnosticReport, index)
		if mostCommonByte == '1' {
			gamaRateStr += "1"
			epsilonRateStr += "0"
		} else {
			gamaRateStr += "0"
			epsilonRateStr += "1"
		}
	}

	gamaRate, err := strconv.ParseInt(gamaRateStr, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilonRate, err := strconv.ParseInt(epsilonRateStr, 2, 64)
	if err != nil {
		panic(err)
	}

	return gamaRate * epsilonRate
}

func calculateOxygenGeneratorRating(diagnosticReport []string) int64 {
	keepers := diagnosticReport

	for index := range diagnosticReport[0] {
		mostCommonByte := mostCommonByteInPosition(keepers, index)

		var newKeepers []string
		for _, k := range keepers {
			if rune(k[index]) == mostCommonByte {
				newKeepers = append(newKeepers, k)
			}
		}

		keepers = newKeepers
		if len(keepers) == 1 {
			break
		}
	}

	oxygenGeneratorRating, err := strconv.ParseInt(keepers[0], 2, 64)
	if err != nil {
		panic(err)
	}

	return oxygenGeneratorRating
}

func calculateCO2ScrubberRating(diagnosticReport []string) int64 {
	keepers := diagnosticReport

	for index := range diagnosticReport[0] {
		mostCommonByte := mostCommonByteInPosition(keepers, index)
		lessCommonByte := '1'
		if mostCommonByte == '1' {
			lessCommonByte = '0'
		}

		var newKeepers []string
		for _, k := range keepers {
			if rune(k[index]) == lessCommonByte {
				newKeepers = append(newKeepers, k)
			}
		}

		keepers = newKeepers
		if len(keepers) == 1 {
			break
		}
	}

	co2ScrubberRating, err := strconv.ParseInt(keepers[0], 2, 64)
	if err != nil {
		panic(err)
	}

	return co2ScrubberRating
}

func CalculateLifeSupportRating(diagnosticReport []string) int64 {
	return calculateOxygenGeneratorRating(diagnosticReport) * calculateCO2ScrubberRating(diagnosticReport)
}

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error occurred: %v", err)
		}
	}()

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(file)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	diagnosticReport, err := pkg.ReadLines(scanner)
	if err != nil {
		panic(err)
	}

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", CalculatePowerConsumption(diagnosticReport)) },
		func() string { return fmt.Sprintf("%v", CalculateLifeSupportRating(diagnosticReport)) },
	)
}
