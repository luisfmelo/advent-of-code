package main

import (
	"bufio"
	"fmt"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
	"log"
	"os"
)

const inputPath = "2021/6/input.txt"

func CalculateNumberOfLanternFishesAfterNDays(lanternFishes []int, days int) int {
	mapLanternFishes := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0}
	for _, fish := range lanternFishes {
		mapLanternFishes[fish]++
	}

	for day := 0; day < days; day++ {
		mapLanternFishes = map[int]int{
			0: mapLanternFishes[1],
			1: mapLanternFishes[2],
			2: mapLanternFishes[3],
			3: mapLanternFishes[4],
			4: mapLanternFishes[5],
			5: mapLanternFishes[6],
			6: mapLanternFishes[7] + mapLanternFishes[0],
			7: mapLanternFishes[8],
			8: mapLanternFishes[0],
		}
	}

	var numberOfFishes int
	for _, n := range mapLanternFishes {
		numberOfFishes += n
	}

	return numberOfFishes
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

	lanternFishes, err := pkg.ReadIntegersInLine(scanner, ",")
	pkg.PanicErr(err)

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", CalculateNumberOfLanternFishesAfterNDays(lanternFishes, 80)) },
		func() string {
			return fmt.Sprintf("%v", CalculateNumberOfLanternFishesAfterNDays(lanternFishes, 256))
		},
	)
}
