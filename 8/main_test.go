package main

import (
	"testing"
)

func TestGetDigit(t *testing.T) {
	testCases := []struct {
		s              string
		expectedOutput int
	}{
		{s: "fdgacbe", expectedOutput: 8},
		{s: "gcbe", expectedOutput: 4},
		{s: "dgebacf", expectedOutput: 8},
		{s: "gc", expectedOutput: 1},
		{s: "cbg", expectedOutput: 7},
	}

	for id, tc := range testCases {
		actualResult := getDigit(tc.s, defaultWireConfig)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCount1s4s7s8s(t *testing.T) {
	testCases := []struct {
		input          []input
		expectedOutput int
	}{
		{
			input: []input{
				{signals: []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}, output: []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}},
				{signals: []string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"}, output: []string{"fcgedb", "cgb", "dgebacf", "gc"}},
				{signals: []string{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"}, output: []string{"cg", "cg", "fdcagb", "cbg"}},
				{signals: []string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"}, output: []string{"efabcd", "cedba", "gadfec", "cb"}},
				{signals: []string{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga"}, output: []string{"gecf", "egdcabf", "bgf", "bfgea"}},
				{signals: []string{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf"}, output: []string{"gebdcfa", "ecba", "ca", "fadegcb"}},
				{signals: []string{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf"}, output: []string{"cefg", "dcbef", "fcge", "gbcadfe"}},
				{signals: []string{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"}, output: []string{"ed", "bcgafe", "cdgba", "cbgef"}},
				{signals: []string{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg"}, output: []string{"gbdfcae", "bgc", "cg", "cgb"}},
				{signals: []string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"}, output: []string{"fgae", "cfgab", "fg", "bagce"}},
			},
			expectedOutput: 26,
		},
	}

	for id, tc := range testCases {
		actualResult := Count1s4s7s8s(tc.input)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestGetWireConfiguration(t *testing.T) {
	testCases := []struct {
		signals        []string
		expectedOutput wireConfiguration
	}{
		{
			signals:        []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
			expectedOutput: wireConfiguration{wireA: 'd', wireB: 'e', wireC: 'a', wireD: 'f', wireE: 'g', wireF: 'b', wireG: 'c'},
		},
	}

	for id, tc := range testCases {
		actualResult := getWireConfiguration(tc.signals)
		for _, wire := range []wireCode{wireA, wireB, wireC, wireD, wireE, wireF, wireG} {
			if actualResult[wire] != tc.expectedOutput[wire] {
				t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult[wire], tc.expectedOutput[wire])
			}
		}
	}
}

func TestSumOutputValues(t *testing.T) {
	testCases := []struct {
		input          []input
		expectedOutput int
	}{
		{
			input: []input{
				{signals: []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}, output: []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"}},
			},
			expectedOutput: 5353,
		},
		{
			input: []input{
				{signals: []string{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"}, output: []string{"ed", "bcgafe", "cdgba", "cbgef"}},
			},
			expectedOutput: 1625,
		},
		{
			input: []input{
				{signals: []string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"}, output: []string{"fgae", "cfgab", "fg", "bagce"}},
			},
			expectedOutput: 4315,
		},
		{
			input: []input{
				{signals: []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}, output: []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}},
				{signals: []string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"}, output: []string{"fcgedb", "cgb", "dgebacf", "gc"}},
				{signals: []string{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"}, output: []string{"cg", "cg", "fdcagb", "cbg"}},
				{signals: []string{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"}, output: []string{"efabcd", "cedba", "gadfec", "cb"}},
				{signals: []string{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga"}, output: []string{"gecf", "egdcabf", "bgf", "bfgea"}},
				{signals: []string{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf"}, output: []string{"gebdcfa", "ecba", "ca", "fadegcb"}},
				{signals: []string{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf"}, output: []string{"cefg", "dcbef", "fcge", "gbcadfe"}},
				{signals: []string{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"}, output: []string{"ed", "bcgafe", "cdgba", "cbgef"}},
				{signals: []string{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg"}, output: []string{"gbdfcae", "bgc", "cg", "cgb"}},
				{signals: []string{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"}, output: []string{"fgae", "cfgab", "fg", "bagce"}},
			},
			expectedOutput: 61229,
		},
	}

	for id, tc := range testCases {
		actualResult := SumOutputValues(tc.input)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
