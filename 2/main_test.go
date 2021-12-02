package main

import (
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
	"testing"
)

func TestMultiplyFinalPosition(t *testing.T) {
	type TestCase struct {
		commands       []types.Command
		expectedOutput int
	}
	testCases := []TestCase{
		{
			commands: []types.Command{
				{Direction: types.Forward, Units: 5},
				{Direction: types.Down, Units: 5},
				{Direction: types.Forward, Units: 8},
				{Direction: types.Up, Units: 3},
				{Direction: types.Down, Units: 8},
				{Direction: types.Forward, Units: 2},
			},
			expectedOutput: 150,
		},
	}

	for id, tc := range testCases {
		actualResult := MultiplyFinalPosition(tc.commands)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestWithAimMultiplyFinalPosition(t *testing.T) {
	type TestCase struct {
		commands       []types.Command
		expectedOutput int
	}
	testCases := []TestCase{
		{
			commands: []types.Command{
				{Direction: types.Forward, Units: 5},
				{Direction: types.Down, Units: 5},
				{Direction: types.Forward, Units: 8},
				{Direction: types.Up, Units: 3},
				{Direction: types.Down, Units: 8},
				{Direction: types.Forward, Units: 2},
			},
			expectedOutput: 900,
		},
	}

	for id, tc := range testCases {
		actualResult := WithAimMultiplyFinalPosition(tc.commands)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
