package main

import (
	"bufio"
	"fmt"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
	"log"
	"math"
	"os"
	"strings"
)

// Original wiring
//  0:      1:      2:      3:      4:
//  aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
//  ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
//  gggg    ....    gggg    gggg    ....
//
//   5:      6:      7:      8:      9:
//  aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
//  dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
//  gggg    gggg    ....    gggg    gggg

const inputPath = "2021/8/input.txt"

type wireCode rune

// wireConfiguration maps actual wiring to the original wiring.
type wireConfiguration map[wireCode]wireCode

const (
	wireA wireCode = 'a'
	wireB wireCode = 'b'
	wireC wireCode = 'c'
	wireD wireCode = 'd'
	wireE wireCode = 'e'
	wireF wireCode = 'f'
	wireG wireCode = 'g'
)

var defaultWireConfig = wireConfiguration{wireA: 'a', wireB: 'b', wireC: 'c', wireD: 'd', wireE: 'e', wireF: 'f', wireG: 'g'}

type input struct {
	signals []string
	output  []string
}

func getDigit(s string, wireConfig wireConfiguration) int {
	switch len(s) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 5:
		var hasWireB, hasWireF bool
		for _, r := range s {
			if r == rune(wireConfig[wireB]) {
				hasWireB = true
			} else if r == rune(wireConfig[wireF]) {
				hasWireF = true
			}
		}
		if hasWireB {
			return 5
		} else if hasWireF {
			return 3
		} else {
			return 2
		}
	case 6:
		missingWireE, missingWireC, missingWireD := true, true, true
		for _, r := range s {
			switch r {
			case rune(wireConfig[wireC]):
				missingWireC = false
			case rune(wireConfig[wireD]):
				missingWireD = false
			case rune(wireConfig[wireE]):
				missingWireE = false
			}
		}
		if missingWireC {
			return 6
		}
		if missingWireD {
			return 0
		}
		if missingWireE {
			return 9
		}
	case 7:
		return 8
	}

	return -1
}

func Count1s4s7s8s(in []input) int {
	var numberOf1s4s7s8s int
	for _, l := range in {
		for _, outputDigitCode := range l.output {
			if d := getDigit(outputDigitCode, defaultWireConfig); d == 1 || d == 4 || d == 7 || d == 8 {
				numberOf1s4s7s8s++
			}
		}
	}

	return numberOf1s4s7s8s
}

func getWireConfiguration(signals []string) wireConfiguration {
	var one, four string
	mapAppearances := map[wireCode]int{wireA: 0, wireB: 0, wireC: 0, wireD: 0, wireE: 0, wireF: 0, wireG: 0}

	// process signals
	for _, signal := range signals {
		if len(signal) == 2 {
			one = signal
		} else if len(signal) == 4 {
			four = signal
		}
		for _, r := range signal {
			mapAppearances[wireCode(r)]++
		}
	}

	// discover wire mapping
	wireConfig := wireConfiguration{}
	for wire, appearances := range mapAppearances {
		switch appearances {
		case 4:
			wireConfig[wireE] = wire
		case 6:
			wireConfig[wireB] = wire
		case 9:
			wireConfig[wireF] = wire
		case 8:
			if strings.Contains(one, string(wire)) {
				wireConfig[wireC] = wire
			} else {
				wireConfig[wireA] = wire
			}
		case 7:
			if strings.Contains(four, string(wire)) {
				wireConfig[wireD] = wire
			} else {
				wireConfig[wireG] = wire
			}
		}
	}

	return wireConfig
}

func getOutputNumber(output []string, wireConfig wireConfiguration) int {
	var outputNumber int
	for index, outputDigitStr := range output {
		outputNumber += getDigit(outputDigitStr, wireConfig) * int(math.Pow(10, float64(3-index)))
	}

	return outputNumber
}

func SumOutputValues(in []input) int {
	var sum int
	for _, l := range in {
		wireConfig := getWireConfiguration(l.signals)
		sum += getOutputNumber(l.output, wireConfig)
	}

	return sum
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

	lines, err := pkg.ReadLines(scanner)
	pkg.PanicErr(err)

	var in []input
	for _, line := range lines {
		sep := strings.Split(line, " | ")

		in = append(in, input{
			signals: strings.Split(sep[0], " "),
			output:  strings.Split(sep[1], " "),
		})
	}

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", Count1s4s7s8s(in)) },
		func() string { return fmt.Sprintf("%v", SumOutputValues(in)) },
	)
}
