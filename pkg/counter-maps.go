package pkg

import "math"

type PairOccurrences struct {
	m map[string]int
}

func NewPairOccurrences() *PairOccurrences {
	return &PairOccurrences{m: map[string]int{}}
}

func (po *PairOccurrences) AddPairOccurrences(pair string, occurrences int) {
	if _, exists := po.m[pair]; !exists {
		po.m[pair] = 0
	}
	po.m[pair] += occurrences
}

func (po *PairOccurrences) GetMapping() map[string]int {
	return po.m
}

type RuneOccurrences struct {
	m map[rune]int
}

func NewRuneOccurrences() *RuneOccurrences {
	return &RuneOccurrences{m: map[rune]int{}}
}

func (po *RuneOccurrences) AddRuneOccurrences(r rune, times int) {
	if _, exists := po.m[r]; !exists {
		po.m[r] = 0
	}
	po.m[r] += times
}

func (po *RuneOccurrences) GetQuantitiesOfMostAndLeastCommonRunes() (mostCommon int, leastCommon int) {
	mostCommon = 0
	leastCommon = math.MaxInt

	for _, occurrences := range po.m {
		if occurrences > mostCommon {
			mostCommon = occurrences
		}
		if occurrences < leastCommon {
			leastCommon = occurrences
		}
	}

	return mostCommon, leastCommon
}
