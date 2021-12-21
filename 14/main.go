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

const inputPath = "14/input.txt"

func GetIndexesNeededToAppend(polymerTemplate string, pairInsertionRules map[string]string) []int {
	var idx []int
	for i := 1; i < len(polymerTemplate); i++ {
		pair := polymerTemplate[i-1 : i+1]
		if _, exists := pairInsertionRules[pair]; exists {
			idx = append(idx, i)
		}
	}

	return idx
}

func ExecuteStep(polymer string, pairInsertionRules map[string]string) string {
	var ptr int
	var newPolymer string
	for _, indexToAppend := range GetIndexesNeededToAppend(polymer, pairInsertionRules) {
		pair := polymer[indexToAppend-1 : indexToAppend+1]
		newPolymer = newPolymer + polymer[ptr:indexToAppend] + pairInsertionRules[pair]
		ptr = indexToAppend
	}

	newPolymer = newPolymer + string(polymer[len(polymer)-1])

	return newPolymer
}

func SubtractQuantityOfMostCommonWithQuantityOfLeastCommon(polymerTemplate string, pairInsertionRules map[string]string, nSteps int) int {
	polymer := polymerTemplate
	for i := 0; i < nSteps; i++ {
		polymer = ExecuteStep(polymer, pairInsertionRules)
	}

	m := map[rune]int{}
	for _, r := range polymer {
		_, exists := m[r]
		if !exists {
			m[r] = 0
		}
		m[r]++
	}

	var (
		mostCommon  = 0
		leastCommon = math.MaxInt
	)
	for _, appearances := range m {
		if appearances > mostCommon {
			mostCommon = appearances
		}
		if appearances < leastCommon {
			leastCommon = appearances
		}
	}

	return mostCommon - leastCommon
}

func GetPairOccurrencesInString(s string) *pkg.PairOccurrences {
	pairOccurrences := pkg.NewPairOccurrences()
	for i := 1; i < len(s); i++ {
		pair := s[i-1 : i+1]
		pairOccurrences.AddPairOccurrences(pair, 1)
	}

	return pairOccurrences
}

func GetRuneOccurrencesBasedOnPairOccurrences(pairOccurrences *pkg.PairOccurrences, lastRune rune) *pkg.RuneOccurrences {
	runeOccurrences := pkg.NewRuneOccurrences()
	for pair, occurrences := range pairOccurrences.GetMapping() {
		// only count the first rune. Otherwise, will duplicate
		runeOccurrences.AddRuneOccurrences(rune(pair[0]), occurrences)
	}

	runeOccurrences.AddRuneOccurrences(lastRune, 1)

	return runeOccurrences
}

func SubtractQuantityOfMostCommonWithQuantityOfLeastCommonSmartely(polymerTemplate string, pairInsertionRules map[string]string, nSteps int) int {
	pairOccurrences := GetPairOccurrencesInString(polymerTemplate)
	for i := 0; i < nSteps; i++ {
		newPairOccurrences := pkg.NewPairOccurrences()
		for pair, occurrences := range pairOccurrences.GetMapping() {
			if insertionRune, ruleApplies := pairInsertionRules[pair]; ruleApplies {
				newPair1 := string(pair[0]) + insertionRune
				newPair2 := insertionRune + string(pair[1])

				newPairOccurrences.AddPairOccurrences(newPair1, occurrences)
				newPairOccurrences.AddPairOccurrences(newPair2, occurrences)
			}
		}

		pairOccurrences = newPairOccurrences
	}

	lastRune := rune(polymerTemplate[len(polymerTemplate)-1])
	runeOccurrences := GetRuneOccurrencesBasedOnPairOccurrences(pairOccurrences, lastRune)
	mostCommon, leastCommon := runeOccurrences.GetQuantitiesOfMostAndLeastCommonRunes()

	return mostCommon - leastCommon
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

	pairInsertionRules := map[string]string{}
	polymerTemplate := fileLines[0]
	for _, line := range fileLines[2:] {
		splitter := strings.Split(line, " -> ")
		pairInsertionRules[splitter[0]] = splitter[1]
	}

	pkg.RunWithTime(
		func() string {
			return fmt.Sprintf("%v", SubtractQuantityOfMostCommonWithQuantityOfLeastCommon(polymerTemplate, pairInsertionRules, 10))
		},
		func() string {
			return fmt.Sprintf("%v", SubtractQuantityOfMostCommonWithQuantityOfLeastCommonSmartely(polymerTemplate, pairInsertionRules, 40))
		},
	)
}
