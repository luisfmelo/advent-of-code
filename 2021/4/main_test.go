package main

import (
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
	"testing"
)

func TestCalculateScoreOfTheWinningBoard(t *testing.T) {
	type TestCase struct {
		numbersDrew    []int
		bingoBallots   []types.Matrix
		expectedOutput int
	}
	testCases := []TestCase{
		{
			numbersDrew: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
			bingoBallots: []types.Matrix{
				{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				{
					{3, 15, 0, 2, 22},
					{9, 18, 13, 17, 5},
					{19, 8, 7, 25, 23},
					{20, 11, 10, 24, 4},
					{14, 21, 16, 12, 6},
				},
				{
					{14, 21, 17, 24, 4},
					{10, 16, 15, 9, 19},
					{18, 8, 23, 26, 20},
					{22, 11, 13, 6, 5},
					{2, 0, 12, 3, 7},
				},
			},
			expectedOutput: 4512,
		},
	}

	for id, tc := range testCases {
		actualResult := CalculateScoreOfTheWinningBoard(tc.numbersDrew, tc.bingoBallots)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCalculateScoreOfTheLatestToWinBoard(t *testing.T) {
	type TestCase struct {
		numbersDrew    []int
		bingoBallots   []types.Matrix
		expectedOutput int
	}
	testCases := []TestCase{
		{
			numbersDrew: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
			bingoBallots: []types.Matrix{
				{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				{
					{3, 15, 0, 2, 22},
					{9, 18, 13, 17, 5},
					{19, 8, 7, 25, 23},
					{20, 11, 10, 24, 4},
					{14, 21, 16, 12, 6},
				},
				{
					{14, 21, 17, 24, 4},
					{10, 16, 15, 9, 19},
					{18, 8, 23, 26, 20},
					{22, 11, 13, 6, 5},
					{2, 0, 12, 3, 7},
				},
			},
			expectedOutput: 1924,
		},
		{
			numbersDrew: []int{1,76,38,96,62,41,27,33,4,2,94,15,89,25,66,14,30,0,71,21,48,44,87,73,60,50,77,45,29,18,5,99,65,16,93,95,37,3,52,32,46,80,98,63,92,24,35,55,12,81,51,17,70,78,61,91,54,8,72,40,74,68,75,67,39,64,10,53,9,31,6,7,47,42,90,20,19,36,22,43,58,28,79,86,57,49,83,84,97,11,85,26,69,23,59,82,88,34,56,13},
			bingoBallots: []types.Matrix{
				{
					{37, 68, 86, 44, 78},
					{87, 67, 77, 70, 60},
					{45, 34, 27, 15, 47},
					{12, 21, 13, 55, 26},
					{81, 41, 63, 40, 74},
				},
			},
			expectedOutput: 48843,
		},
	}

	for id, tc := range testCases {
		actualResult := CalculateScoreOfTheLatestToWinBoard(tc.numbersDrew, tc.bingoBallots)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCalculateScoreOfBoard(t *testing.T) {
	type TestCase struct {
		board          types.Matrix
		lastNumberDrew int
		expectedOutput int
	}
	testCases := []TestCase{
		{
			board: types.Matrix{
				{cross, cross, cross, cross, cross},
				{1, 2, 3, cross, 4},    // 10
				{cross, 5, 10, 15, 20}, // 50
				{cross, 5, 10, 15, 20}, // 50
				{cross, 5, 10, 15, 20}, // 50
			},
			lastNumberDrew: 10,
			expectedOutput: 1600,
		},
		{
			board: types.Matrix{
				{cross, cross, cross, cross, cross},
				{0, 0, 0, cross, 0}, // 10
				{cross, 0, 0, 0, 0}, // 50
				{cross, 0, 0, 0, 0}, // 50
				{cross, 0, 0, 0, 0}, // 50
			},
			lastNumberDrew: 11,
			expectedOutput: 0,
		},
		{
			board: types.Matrix{
				{cross, cross, cross, cross, cross},
				{0, 0, 0, cross, 0}, // 10
				{cross, 1, 0, 0, 0}, // 50
				{cross, 0, 0, 1, 0}, // 50
				{cross, 0, 0, 0, 0}, // 50
			},
			lastNumberDrew: 17,
			expectedOutput: 34,
		},
		{
			board: types.Matrix{
				{-1, 91, 37, -1, 98},
				{68, -1, 34, -1, 43},
				{75, -1, 67, -1, 69},
				{81, 47, 58, -1, 93},
				{88, 92, 42, -1, 54},
			},
			lastNumberDrew: 5,
			expectedOutput: 1137 * 5,
		},
	}

	for id, tc := range testCases {
		actualResult := calculateScoreOfBoard(tc.board, tc.lastNumberDrew)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCheckRowWin(t *testing.T) {
	type TestCase struct {
		board          types.Matrix
		rowIndex       int
		expectedOutput bool
	}
	testCases := []TestCase{
		{
			board: types.Matrix{
				{1, 2, 3, cross, 4},
				{cross, cross, cross, cross, cross},
				{cross, 5, 10, 15, 20},
				{cross, 5, 10, 15, 20},
				{cross, 5, 10, 15, 20},
			},
			rowIndex:       1,
			expectedOutput: true,
		},
		{
			board: types.Matrix{
				{cross, 0, 0, 0, 0},
				{0, cross, 0, 0, 0},
				{0, 0, 0, cross, 0},
				{0, 0, 0, cross, 0},
				{0, 0, 0, 0, cross},
			},
			rowIndex:       3,
			expectedOutput: false,
		},
		{
			board: types.Matrix{
				{cross, 0, 0, 0, 0},
				{0, cross, 0, 0, 0},
				{0, 0, 0, cross, 0},
				{cross, cross, cross, cross, 0},
				{0, 0, 0, 0, cross},
			},
			rowIndex:       3,
			expectedOutput: false,
		},
	}

	for id, tc := range testCases {
		actualResult := checkRowWin(tc.board, tc.rowIndex)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCheckColumnWin(t *testing.T) {
	type TestCase struct {
		board          types.Matrix
		colIndex       int
		expectedOutput bool
	}
	testCases := []TestCase{
		{
			board: types.Matrix{
				{1, 2, 3, cross, 4},
				{cross, cross, cross, cross, cross},
				{cross, 5, 10, 15, 20},
				{cross, 5, 10, 15, 20},
				{cross, 5, 10, 15, 20},
			},
			colIndex:       0,
			expectedOutput: false,
		},
		{
			board: types.Matrix{
				{cross, 0, 0, 0, 0},
				{0, cross, 0, 0, 0},
				{0, 0, 0, cross, 0},
				{0, 0, 0, cross, 0},
				{0, 0, 0, 0, cross},
			},
			colIndex:       4,
			expectedOutput: false,
		},
		{
			board: types.Matrix{
				{cross, 0, 0, 0, 0},
				{cross, cross, 0, 0, 0},
				{cross, 0, 0, cross, 0},
				{cross, cross, cross, cross, 0},
				{cross, 0, 0, 0, cross},
			},
			colIndex:       0,
			expectedOutput: true,
		},
	}

	for id, tc := range testCases {
		actualResult := checkColumnWin(tc.board, tc.colIndex)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
