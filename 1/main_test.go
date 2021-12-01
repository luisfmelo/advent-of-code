package main

import (
	"testing"
)

func TestCountMeasurementIncreases(t *testing.T) {
	type TestCase struct {
		measurements   []int
		expectedOutput int
	}
	testCases := []TestCase{
		{
			measurements:   []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			expectedOutput: 7,
		},
	}

	for id, tc := range testCases {
		actualResult := CountMeasurementIncreases(tc.measurements)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCount3MeasurementWindowIncreases(t *testing.T) {
	type TestCase struct {
		measurements   []int
		expectedOutput int
	}
	testCases := []TestCase{
		{
			measurements:   []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			expectedOutput: 5,
		},
	}

	for id, tc := range testCases {
		actualResult := Count3MeasurementWindowIncreases(tc.measurements)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
