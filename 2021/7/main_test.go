package main

import (
	"testing"
)

func TestCalculateLeastAmountOfFuelPart1(t *testing.T) {
	testCases := []struct {
		crabPositions  []int
		expectedOutput int
	}{
		{
			crabPositions:  []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			expectedOutput: 37,
		},
	}

	for id, tc := range testCases {
		actualResult := CalculateLeastAmountOfFuelWhenBurnedAtConstantRate(tc.crabPositions)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCalculateLeastAmountOfFuelPart2(t *testing.T) {
	testCases := []struct {
		crabPositions  []int
		expectedOutput int
	}{
		{
			crabPositions:  []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			expectedOutput: 168,
		},
	}

	for id, tc := range testCases {
		actualResult := CalculateLeastAmountOfFuelWhenBurnedAtIncreasingRate(tc.crabPositions)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
