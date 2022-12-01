package main

import (
	"testing"
)

func TestMaxFoodCarriedBySingleElf(t *testing.T) {
	type TestCase struct {
		input          [][]int
		expectedOutput int
	}
	testCases := []TestCase{
		{
			input:          [][]int{{1000, 2000, 3000}, {4000}, {5000, 6000}, {7000, 8000, 9000}, {10000}},
			expectedOutput: 24000,
		},
	}

	for id, tc := range testCases {
		actualResult := MaxFoodCarriedBySingleElf(tc.input)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestMaxFoodCarriedByTop3Elfs(t *testing.T) {
	type TestCase struct {
		input          [][]int
		expectedOutput int
	}
	testCases := []TestCase{
		{
			input:          [][]int{{1000, 2000, 3000}, {4000}, {5000, 6000}, {7000, 8000, 9000}, {10000}},
			expectedOutput: 45000,
		},
	}

	for id, tc := range testCases {
		actualResult := MaxFoodCarriedByTop3Elfs(tc.input)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
