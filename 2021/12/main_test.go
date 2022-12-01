package main

import (
	"testing"
)

func TestCountNumberOfPaths(t *testing.T) {
	testCases := []struct {
		inputLines     []string
		expectedOutput int
	}{
		{
			inputLines: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			expectedOutput: 10,
		},
		//{
		//	inputLines: []string{
		//		"dc-end",
		//		"HN-start",
		//		"start-kj",
		//		"dc-start",
		//		"dc-HN",
		//		"LN-dc",
		//		"HN-end",
		//		"kj-sa",
		//		"kj-HN",
		//		"kj-dc",
		//	},
		//	expectedOutput: 19,
		//},
		//{
		//	inputLines: []string{
		//		"fs-end",
		//		"he-DX",
		//		"fs-he",
		//		"start-DX",
		//		"pj-DX",
		//		"end-zg",
		//		"zg-sl",
		//		"zg-pj",
		//		"pj-he",
		//		"RW-he",
		//		"fs-DX",
		//		"pj-RW",
		//		"zg-RW",
		//		"start-pj",
		//		"he-WI",
		//		"zg-he",
		//		"pj-fs",
		//		"start-RW",
		//	},
		//	expectedOutput: 226,
		//},
	}

	for id, tc := range testCases {
		actualResult := CountNumberOfPaths(tc.inputLines)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
