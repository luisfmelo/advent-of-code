package main

import (
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
	"testing"
)

func TestGetPointsInYAxis(t *testing.T) {
	type TestCase struct {
		p1             types.Point
		p2             types.Point
		expectedOutput []types.Point
	}
	testCases := []TestCase{
		{
			p1:             types.Point{X: 0, Y: 9},
			p2:             types.Point{X: 3, Y: 9},
			expectedOutput: []types.Point{{0, 9}, {1, 9}, {2, 9}, {3, 9}},
		},
		{
			p1:             types.Point{X: 3, Y: 9},
			p2:             types.Point{X: 0, Y: 9},
			expectedOutput: []types.Point{{0, 9}, {1, 9}, {2, 9}, {3, 9}},
		},
		{
			p1:             types.Point{X: 3, Y: 9},
			p2:             types.Point{X: 3, Y: 5},
			expectedOutput: []types.Point{},
		},
		{
			p1:             types.Point{X: 2, Y: 2},
			p2:             types.Point{X: 1, Y: 2},
			expectedOutput: []types.Point{{X: 2, Y: 2}, {X: 2, Y: 1}},
		},
	}

	for id, tc := range testCases {
		actualResult := getPointsInXAxis(tc.p1, tc.p2)
		if len(actualResult) != len(tc.expectedOutput) {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestGetPointsInXAxis(t *testing.T) {
	type TestCase struct {
		p1             types.Point
		p2             types.Point
		expectedOutput []types.Point
	}
	testCases := []TestCase{
		{
			p1:             types.Point{X: 9, Y: 0},
			p2:             types.Point{X: 9, Y: 3},
			expectedOutput: []types.Point{{9, 0}, {9, 1}, {9, 2}, {9, 3}},
		},
		{
			p1:             types.Point{X: 9, Y: 3},
			p2:             types.Point{X: 9, Y: 0},
			expectedOutput: []types.Point{{9, 0}, {9, 1}, {9, 2}, {9, 3}},
		},
		{
			p1:             types.Point{X: 9, Y: 5},
			p2:             types.Point{X: 5, Y: 5},
			expectedOutput: []types.Point{},
		},
	}

	for id, tc := range testCases {
		actualResult := getPointsInYAxis(tc.p1, tc.p2)
		if len(actualResult) != len(tc.expectedOutput) {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestTGetPointsInDiagonal(t *testing.T) {
	type TestCase struct {
		p1             types.Point
		p2             types.Point
		expectedOutput []types.Point
	}
	testCases := []TestCase{
		{
			p1:             types.Point{X: 3, Y: 0},
			p2:             types.Point{X: 0, Y: 3},
			expectedOutput: []types.Point{{3, 0}, {2, 1}, {1, 2}, {0, 3}},
		},
	}

	for id, tc := range testCases {
		actualResult := getPointsInDiagonal(tc.p1, tc.p2)
		if len(actualResult) != len(tc.expectedOutput) {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCheckNumberOfOverlaps(t *testing.T) {
	type TestCase struct {
		lines          []types.Line
		withDiagonals  bool
		expectedOutput int
	}
	testCases := []TestCase{
		{
			lines: []types.Line{
				{Start: types.Point{X: 0, Y: 9}, End: types.Point{X: 5, Y: 9}},
				{Start: types.Point{X: 8, Y: 0}, End: types.Point{X: 0, Y: 8}},
				{Start: types.Point{X: 9, Y: 4}, End: types.Point{X: 3, Y: 4}},
				{Start: types.Point{X: 2, Y: 2}, End: types.Point{X: 2, Y: 1}},
				{Start: types.Point{X: 7, Y: 0}, End: types.Point{X: 7, Y: 4}},
				{Start: types.Point{X: 6, Y: 4}, End: types.Point{X: 2, Y: 0}},
				{Start: types.Point{X: 0, Y: 9}, End: types.Point{X: 2, Y: 9}},
				{Start: types.Point{X: 3, Y: 4}, End: types.Point{X: 1, Y: 4}},
				{Start: types.Point{X: 0, Y: 0}, End: types.Point{X: 8, Y: 8}},
				{Start: types.Point{X: 5, Y: 5}, End: types.Point{X: 8, Y: 2}},
			},
			withDiagonals:  false,
			expectedOutput: 5,
		},
		{
			lines: []types.Line{
				{Start: types.Point{X: 0, Y: 9}, End: types.Point{X: 5, Y: 9}},
				{Start: types.Point{X: 8, Y: 0}, End: types.Point{X: 0, Y: 8}},
				{Start: types.Point{X: 9, Y: 4}, End: types.Point{X: 3, Y: 4}},
				{Start: types.Point{X: 2, Y: 2}, End: types.Point{X: 2, Y: 1}},
				{Start: types.Point{X: 7, Y: 0}, End: types.Point{X: 7, Y: 4}},
				{Start: types.Point{X: 6, Y: 4}, End: types.Point{X: 2, Y: 0}},
				{Start: types.Point{X: 0, Y: 9}, End: types.Point{X: 2, Y: 9}},
				{Start: types.Point{X: 3, Y: 4}, End: types.Point{X: 1, Y: 4}},
				{Start: types.Point{X: 0, Y: 0}, End: types.Point{X: 8, Y: 8}},
				{Start: types.Point{X: 5, Y: 5}, End: types.Point{X: 8, Y: 2}},
			},
			withDiagonals:  true,
			expectedOutput: 12,
		},
	}

	for id, tc := range testCases {
		actualResult := CheckNumberOfOverlaps(tc.lines, tc.withDiagonals)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
