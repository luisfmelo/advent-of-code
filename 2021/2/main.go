package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
)

const inputPath = "2021/2/input.txt"

func MultiplyFinalPosition(commands []types.Command) int {
	var hPos, vPos int
	for _, command := range commands {
		switch command.Direction {
		case types.Forward:
			hPos += command.Units
		case types.Down:
			vPos += command.Units
		case types.Up:
			vPos -= command.Units
		}
	}

	return vPos * hPos
}

func WithAimMultiplyFinalPosition(commands []types.Command) int {
	var hPos, vPos, aim int
	for _, command := range commands {
		switch command.Direction {
		case types.Forward:
			hPos += command.Units
			vPos += aim * command.Units
		case types.Down:
			aim += command.Units
		case types.Up:
			aim -= command.Units
		}
	}

	return vPos * hPos
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

	input, err := pkg.ReadLines(scanner)
	if err != nil {
		panic(err)
	}

	var commands []types.Command
	for _, l := range input {
		splitter := strings.Split(l, " ")

		var units int
		units, err = strconv.Atoi(splitter[1])
		if err != nil {
			panic(err)
		}

		commands = append(commands, types.Command{Direction: types.Direction(splitter[0]), Units: units})
	}

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", MultiplyFinalPosition(commands)) },
		func() string { return fmt.Sprintf("%v", WithAimMultiplyFinalPosition(commands)) },
	)
}
