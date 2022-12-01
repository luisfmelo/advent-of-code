package main

import (
	"testing"

	"github.com/luisfmelo/go-advent-of-code-2021/pkg/test"
)

func TestGetIndexesNeededToAppend(t *testing.T) {
	type TestCase struct {
		polymerTemplate    string
		pairInsertionRules map[string]string
		expectedOutput     []int
	}
	testCases := []TestCase{
		{
			polymerTemplate: "NNCB",
			pairInsertionRules: map[string]string{
				"CH": "B",
				"HH": "N",
				"CB": "H",
				"NH": "C",
				"HB": "C",
				"HC": "B",
				"HN": "C",
				"NN": "C",
				"BH": "H",
				"NC": "B",
				"NB": "B",
				"BN": "B",
				"BB": "N",
				"BC": "B",
				"CC": "N",
				"CN": "C",
			},
			expectedOutput: []int{1, 2, 3},
		},
	}

	for id, tc := range testCases {
		actualResult := GetIndexesNeededToAppend(tc.polymerTemplate, tc.pairInsertionRules)
		if !test.IntArrIsEqual(actualResult, tc.expectedOutput) {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestExecuteStep(t *testing.T) {
	type TestCase struct {
		polymer            string
		pairInsertionRules map[string]string
		expectedOutput     string
	}
	testCases := []TestCase{
		{
			polymer: "NNCB",
			pairInsertionRules: map[string]string{
				"CH": "B",
				"HH": "N",
				"CB": "H",
				"NH": "C",
				"HB": "C",
				"HC": "B",
				"HN": "C",
				"NN": "C",
				"BH": "H",
				"NC": "B",
				"NB": "B",
				"BN": "B",
				"BB": "N",
				"BC": "B",
				"CC": "N",
				"CN": "C",
			},
			expectedOutput: "NCNBCHB",
		},
		{
			polymer: "NCNBCHB",
			pairInsertionRules: map[string]string{
				"CH": "B",
				"HH": "N",
				"CB": "H",
				"NH": "C",
				"HB": "C",
				"HC": "B",
				"HN": "C",
				"NN": "C",
				"BH": "H",
				"NC": "B",
				"NB": "B",
				"BN": "B",
				"BB": "N",
				"BC": "B",
				"CC": "N",
				"CN": "C",
			},
			expectedOutput: "NBCCNBBBCBHCB",
		},
		{
			polymer: "NBCCNBBBCBHCB",
			pairInsertionRules: map[string]string{
				"CH": "B",
				"HH": "N",
				"CB": "H",
				"NH": "C",
				"HB": "C",
				"HC": "B",
				"HN": "C",
				"NN": "C",
				"BH": "H",
				"NC": "B",
				"NB": "B",
				"BN": "B",
				"BB": "N",
				"BC": "B",
				"CC": "N",
				"CN": "C",
			},
			expectedOutput: "NBBBCNCCNBBNBNBBCHBHHBCHB",
		},
		{
			polymer: "NBBBCNCCNBBNBNBBCHBHHBCHB",
			pairInsertionRules: map[string]string{
				"CH": "B",
				"HH": "N",
				"CB": "H",
				"NH": "C",
				"HB": "C",
				"HC": "B",
				"HN": "C",
				"NN": "C",
				"BH": "H",
				"NC": "B",
				"NB": "B",
				"BN": "B",
				"BB": "N",
				"BC": "B",
				"CC": "N",
				"CN": "C",
			},
			expectedOutput: "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB",
		},
	}

	for id, tc := range testCases {
		actualResult := ExecuteStep(tc.polymer, tc.pairInsertionRules)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestSubtractQuantityOfMostCommonWithQuantityOfLeastCommon(t *testing.T) {
	type TestCase struct {
		polymer            string
		pairInsertionRules map[string]string
		nSteps             int
		expectedOutput     int
	}
	testCases := []TestCase{
		{
			polymer: "NNCB",
			pairInsertionRules: map[string]string{
				"CH": "B",
				"HH": "N",
				"CB": "H",
				"NH": "C",
				"HB": "C",
				"HC": "B",
				"HN": "C",
				"NN": "C",
				"BH": "H",
				"NC": "B",
				"NB": "B",
				"BN": "B",
				"BB": "N",
				"BC": "B",
				"CC": "N",
				"CN": "C",
			},
			nSteps:         4,
			expectedOutput: 18,
		},
		{
			polymer: "NNCB",
			pairInsertionRules: map[string]string{
				"CH": "B",
				"HH": "N",
				"CB": "H",
				"NH": "C",
				"HB": "C",
				"HC": "B",
				"HN": "C",
				"NN": "C",
				"BH": "H",
				"NC": "B",
				"NB": "B",
				"BN": "B",
				"BB": "N",
				"BC": "B",
				"CC": "N",
				"CN": "C",
			},
			nSteps:         10,
			expectedOutput: 1588,
		},
	}

	for id, tc := range testCases {
		actualResult := SubtractQuantityOfMostCommonWithQuantityOfLeastCommon(tc.polymer, tc.pairInsertionRules, tc.nSteps)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestSubtractQuantityOfMostCommonWithQuantityOfLeastCommonSmartely(t *testing.T) {
	type TestCase struct {
		polymer            string
		pairInsertionRules map[string]string
		nSteps             int
		expectedOutput     int
	}
	testCases := []TestCase{
		{
			polymer: "NNCB",
			pairInsertionRules: map[string]string{
				"CH": "B",
				"HH": "N",
				"CB": "H",
				"NH": "C",
				"HB": "C",
				"HC": "B",
				"HN": "C",
				"NN": "C",
				"BH": "H",
				"NC": "B",
				"NB": "B",
				"BN": "B",
				"BB": "N",
				"BC": "B",
				"CC": "N",
				"CN": "C",
			},
			nSteps:         1,
			expectedOutput: 1,
		},
		{
			polymer: "NNCB",
			pairInsertionRules: map[string]string{
				"CH": "B",
				"HH": "N",
				"CB": "H",
				"NH": "C",
				"HB": "C",
				"HC": "B",
				"HN": "C",
				"NN": "C",
				"BH": "H",
				"NC": "B",
				"NB": "B",
				"BN": "B",
				"BB": "N",
				"BC": "B",
				"CC": "N",
				"CN": "C",
			},
			nSteps:         4,
			expectedOutput: 18,
		},
		{
			polymer: "NNCB",
			pairInsertionRules: map[string]string{
				"CH": "B",
				"HH": "N",
				"CB": "H",
				"NH": "C",
				"HB": "C",
				"HC": "B",
				"HN": "C",
				"NN": "C",
				"BH": "H",
				"NC": "B",
				"NB": "B",
				"BN": "B",
				"BB": "N",
				"BC": "B",
				"CC": "N",
				"CN": "C",
			},
			nSteps:         10,
			expectedOutput: 1588,
		},
		{
			polymer: "NNCB",
			pairInsertionRules: map[string]string{
				"CH": "B",
				"HH": "N",
				"CB": "H",
				"NH": "C",
				"HB": "C",
				"HC": "B",
				"HN": "C",
				"NN": "C",
				"BH": "H",
				"NC": "B",
				"NB": "B",
				"BN": "B",
				"BB": "N",
				"BC": "B",
				"CC": "N",
				"CN": "C",
			},
			nSteps:         40,
			expectedOutput: 2188189693529,
		},
	}

	for id, tc := range testCases {
		actualResult := SubtractQuantityOfMostCommonWithQuantityOfLeastCommonSmartely(tc.polymer, tc.pairInsertionRules, tc.nSteps)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
