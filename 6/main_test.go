package main

import (
	"testing"
)

func TestCalculateNumberOfLanternFishesAfterNDays(t *testing.T) {
	testCases := []struct {
		initialLanternFishes []int
		days                 int
		expectedOutput       int
	}{
		{
			initialLanternFishes: []int{3, 4, 3, 1, 2},
			days:                 18,
			expectedOutput:       26,
		},
		{
			initialLanternFishes: []int{3, 4, 3, 1, 2},
			days:                 80,
			expectedOutput:       5934,
		},
	}

	for id, tc := range testCases {
		actualResult := CalculateNumberOfLanternFishesAfterNDays(tc.initialLanternFishes, tc.days)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCalculateNumberOfLanternFishesAfterNDaysInParallel(t *testing.T) {
	testCases := []struct {
		initialLanternFishes []int
		days                 int
		expectedOutput       int
	}{
		{
			initialLanternFishes: []int{3, 4, 3, 1, 2},
			days:                 18,
			expectedOutput:       26,
		},
		{
			initialLanternFishes: []int{3, 4, 3, 1, 2},
			days:                 80,
			expectedOutput:       5934,
		},
		{
			initialLanternFishes: []int{3, 4, 3, 1, 2},
			days:                 256,
			expectedOutput:       26984457539,
		},
	}

	for id, tc := range testCases {
		actualResult := CalculateNumberOfLanternFishesAfterNDays(tc.initialLanternFishes, tc.days)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
