package main

import (
	"testing"
)

func TestCalculatePowerConsumption(t *testing.T) {
	type TestCase struct {
		diagnosticReport []string
		expectedOutput   int64
	}
	testCases := []TestCase{
		{
			diagnosticReport: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			expectedOutput: 198,
		},
	}

	for id, tc := range testCases {
		actualResult := CalculatePowerConsumption(tc.diagnosticReport)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCalculateOxygenGeneratorRating(t *testing.T) {
	type TestCase struct {
		diagnosticReport []string
		expectedOutput   int64
	}
	testCases := []TestCase{
		{
			diagnosticReport: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			expectedOutput: 23,
		},
	}

	for id, tc := range testCases {
		actualResult := calculateOxygenGeneratorRating(tc.diagnosticReport)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCalculateCO2ScrubberRating(t *testing.T) {
	type TestCase struct {
		diagnosticReport []string
		expectedOutput   int64
	}
	testCases := []TestCase{
		{
			diagnosticReport: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			expectedOutput: 10,
		},
	}

	for id, tc := range testCases {
		actualResult := calculateCO2ScrubberRating(tc.diagnosticReport)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCalculateLifeSupportRating(t *testing.T) {
	type TestCase struct {
		diagnosticReport []string
		expectedOutput   int64
	}
	testCases := []TestCase{
		{
			diagnosticReport: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			expectedOutput: 230,
		},
	}

	for id, tc := range testCases {
		actualResult := CalculateLifeSupportRating(tc.diagnosticReport)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
