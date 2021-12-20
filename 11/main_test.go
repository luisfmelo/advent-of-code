package main

import (
	"testing"

	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
)

func TestCountNumberOfFlashes(t *testing.T) {
	testCases := []struct {
		matrix         types.Matrix
		steps          int
		expectedOutput int
	}{
		{
			matrix: types.Matrix{
				{1, 1, 1, 1, 1},
				{1, 9, 9, 9, 1},
				{1, 9, 1, 9, 1},
				{1, 9, 9, 9, 1},
				{1, 1, 1, 1, 1},
			},
			steps:          1,
			expectedOutput: 9,
		},
		{
			matrix: types.Matrix{
				{1, 1, 1, 1, 1},
				{1, 9, 9, 9, 1},
				{1, 9, 1, 9, 1},
				{1, 9, 9, 9, 1},
				{1, 1, 1, 1, 1},
			},
			steps:          2,
			expectedOutput: 9,
		},
		{
			matrix: types.Matrix{
				{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
				{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
				{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
				{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
				{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
				{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
				{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
				{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
				{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
				{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
			},
			steps:          1,
			expectedOutput: 0,
		},
		{
			matrix: types.Matrix{
				{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
				{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
				{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
				{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
				{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
				{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
				{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
				{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
				{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
				{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
			},
			steps:          2,
			expectedOutput: 35,
		},
		{
			matrix: types.Matrix{
				{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
				{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
				{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
				{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
				{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
				{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
				{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
				{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
				{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
				{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
			},
			steps:          10,
			expectedOutput: 204,
		},
		{
			matrix: types.Matrix{
				{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
				{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
				{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
				{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
				{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
				{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
				{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
				{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
				{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
				{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
			},
			steps:          100,
			expectedOutput: 1656,
		},
	}

	for id, tc := range testCases {
		actualResult := CountNumberOfFlashes(tc.matrix, tc.steps)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCountNumberOfStepsUntilSimultaneousFlashes(t *testing.T) {
	testCases := []struct {
		matrix         types.Matrix
		expectedOutput int
	}{
		{
			matrix: types.Matrix{
				{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
				{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
				{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
				{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
				{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
				{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
				{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
				{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
				{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
				{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
			},
			expectedOutput: 195,
		},
	}

	for id, tc := range testCases {
		actualResult := CountNumberOfStepsUntilSimultaneousFlashes(tc.matrix)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
