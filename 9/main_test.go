package main

import (
	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"

	"testing"
)

func TestGetLavaTunnelLowLevels(t *testing.T) {
	testCases := []struct {
		lavaTunnelHeights types.Matrix
		expectedOutput    []types.MatrixElement
	}{
		{
			lavaTunnelHeights: types.Matrix{
				{1, 2, 3},
				{3, 9, 8},
			},
			expectedOutput: []types.MatrixElement{{LineIndex: 0, ColumnIndex: 0}},
		},
		{
			lavaTunnelHeights: types.Matrix{
				{3, 3, 3},
				{3, 2, 3},
				{3, 3, 3},
			},
			expectedOutput: []types.MatrixElement{{LineIndex: 1, ColumnIndex: 1}},
		},
		{
			lavaTunnelHeights: types.Matrix{
				{3, 3, 3},
				{3, 2, 3},
			},
			expectedOutput: []types.MatrixElement{{LineIndex: 1, ColumnIndex: 1}},
		},
		{
			lavaTunnelHeights: types.Matrix{
				{3, 3, 3},
				{3, 3, 2},
			},
			expectedOutput: []types.MatrixElement{{LineIndex: 1, ColumnIndex: 2}},
		},
		{
			lavaTunnelHeights: types.Matrix{
				{3, 3, 3},
				{2, 3, 2},
			},
			expectedOutput: []types.MatrixElement{{LineIndex: 1, ColumnIndex: 0}, {LineIndex: 1, ColumnIndex: 2}},
		},
		{
			lavaTunnelHeights: types.Matrix{
				{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
				{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
				{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
				{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
				{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
			},
			expectedOutput: []types.MatrixElement{{LineIndex: 0, ColumnIndex: 1}, {LineIndex: 0, ColumnIndex: 9}, {LineIndex: 2, ColumnIndex: 2}, {LineIndex: 4, ColumnIndex: 6}},
		},
	}

	for id, tc := range testCases {
		actualResult := getLavaTunnelLowLevels(tc.lavaTunnelHeights)
		if !pkg.MatrixElementArrIsEqual(actualResult, tc.expectedOutput) {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestSumRiskLevelsOfLavaTunnel(t *testing.T) {
	testCases := []struct {
		lavaTunnelHeights types.Matrix
		expectedOutput    int
	}{
		{
			lavaTunnelHeights: types.Matrix{
				{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
				{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
				{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
				{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
				{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
			},
			expectedOutput: 15,
		},
	}

	for id, tc := range testCases {
		actualResult := SumRiskLevelsOfLavaTunnel(tc.lavaTunnelHeights)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestBasinSize(t *testing.T) {
	testCases := []struct {
		lavaTunnelHeights types.Matrix
		basinCenter       types.MatrixElement
		expectedOutput    int
	}{
		{
			lavaTunnelHeights: types.Matrix{
				{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
				{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
				{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
				{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
				{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
			},
			basinCenter:    types.MatrixElement{LineIndex: 0, ColumnIndex: 1},
			expectedOutput: 3,
		},
		{
			lavaTunnelHeights: types.Matrix{
				{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
				{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
				{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
				{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
				{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
			},
			basinCenter:    types.MatrixElement{LineIndex: 0, ColumnIndex: 9},
			expectedOutput: 9,
		},
		{
			lavaTunnelHeights: types.Matrix{
				{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
				{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
				{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
				{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
				{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
			},
			basinCenter:    types.MatrixElement{LineIndex: 2, ColumnIndex: 2},
			expectedOutput: 14,
		},
		{
			lavaTunnelHeights: types.Matrix{
				{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
				{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
				{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
				{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
				{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
			},
			basinCenter:    types.MatrixElement{LineIndex: 4, ColumnIndex: 6},
			expectedOutput: 9,
		},
	}

	for id, tc := range testCases {
		actualResult := GetBasinSize(tc.lavaTunnelHeights, tc.basinCenter)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestMultiply3GreatestBasins(t *testing.T) {
	testCases := []struct {
		lavaTunnelHeights types.Matrix
		expectedOutput    int
	}{
		{
			lavaTunnelHeights: types.Matrix{
				{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
				{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
				{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
				{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
				{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
			},
			expectedOutput: 1134,
		},
	}

	for id, tc := range testCases {
		actualResult := Multiply3GreatestBasins(tc.lavaTunnelHeights)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
