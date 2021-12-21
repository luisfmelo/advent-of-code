package main

import (
	"testing"

	"github.com/luisfmelo/go-advent-of-code-2021/pkg/test"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
)

func TestCountDots(t *testing.T) {
	type TestCase struct {
		m              types.Matrix
		expectedOutput int
	}
	testCases := []TestCase{
		{
			m: types.Matrix{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			expectedOutput: 16,
		},
	}

	for id, tc := range testCases {
		actualResult := countDots(tc.m)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestFold(t *testing.T) {
	type TestCase struct {
		m              types.Matrix
		instruction    Instruction
		expectedOutput types.Matrix
	}
	testCases := []TestCase{
		{
			m: types.Matrix{
				{0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			instruction: Instruction{
				fType:  foldHorizontal,
				cutoff: 7,
			},
			expectedOutput: types.Matrix{
				{1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0},
				{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
				{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			m: types.Matrix{
				{1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0},
				{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
				{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			instruction: Instruction{
				fType:  foldVertical,
				cutoff: 5,
			},
			expectedOutput: types.Matrix{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
		},
	}

	for id, tc := range testCases {
		actualResult := fold(tc.m, tc.instruction)
		if !test.MatrixIsEqual(actualResult, tc.expectedOutput) {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCountNumberOfDots(t *testing.T) {
	type TestCase struct {
		dots           []types.Point
		instructions   []Instruction
		expectedOutput int
	}
	testCases := []TestCase{
		{
			dots: []types.Point{
				{X: 6, Y: 10},
				{X: 0, Y: 14},
				{X: 9, Y: 10},
				{X: 0, Y: 3},
				{X: 10, Y: 4},
				{X: 4, Y: 11},
				{X: 6, Y: 0},
				{X: 6, Y: 12},
				{X: 4, Y: 1},
				{X: 0, Y: 13},
				{X: 10, Y: 12},
				{X: 3, Y: 4},
				{X: 3, Y: 0},
				{X: 8, Y: 4},
				{X: 1, Y: 10},
				{X: 2, Y: 14},
				{X: 8, Y: 10},
				{X: 9, Y: 0},
			},
			instructions: []Instruction{
				NewInstruction("fold along y=7"),
			},
			expectedOutput: 17,
		},
		{
			dots: []types.Point{
				{X: 6, Y: 10},
				{X: 0, Y: 14},
				{X: 9, Y: 10},
				{X: 0, Y: 3},
				{X: 10, Y: 4},
				{X: 4, Y: 11},
				{X: 6, Y: 0},
				{X: 6, Y: 12},
				{X: 4, Y: 1},
				{X: 0, Y: 13},
				{X: 10, Y: 12},
				{X: 3, Y: 4},
				{X: 3, Y: 0},
				{X: 8, Y: 4},
				{X: 1, Y: 10},
				{X: 2, Y: 14},
				{X: 8, Y: 10},
				{X: 9, Y: 0},
			},
			instructions: []Instruction{
				NewInstruction("fold along y=7"),
				NewInstruction("fold along x=5"),
			},
			expectedOutput: 16,
		},
	}

	for id, tc := range testCases {
		actualResult := CountNumberOfDots(tc.dots, tc.instructions...)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
