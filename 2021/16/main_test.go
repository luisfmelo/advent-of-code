package main

import (
	"testing"
)

func TestProcessLiteralValue(t *testing.T) {
	type TestCase struct {
		literalValueString string
		expectedOutput     int
	}
	testCases := []TestCase{
		{
			literalValueString: "00001",
			expectedOutput:     1,
		},
		{
			literalValueString: "00011",
			expectedOutput:     3,
		},
		{
			literalValueString: "101111111000101",
			expectedOutput:     2021,
		},
	}

	for id, tc := range testCases {
		actualResult, _ := ProcessLiteralValue(tc.literalValueString)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestSumVersionsOfPackets(t *testing.T) {
	type TestCase struct {
		input          string
		expectedOutput int
	}
	testCases := []TestCase{
		{
			input:          "EE00D40C823060",
			expectedOutput: 14,
		},
		{
			input:          "8A004A801A8002F478",
			expectedOutput: 16,
		},
		{
			input:          "620080001611562C8802118E34",
			expectedOutput: 12,
		},
		{
			input:          "C0015000016115A2E0802F182340",
			expectedOutput: 23,
		},
		{
			input:          "A0016C880162017C3686B18A3D4780",
			expectedOutput: 31,
		},
	}

	for id, tc := range testCases {
		actualResult := SumVersionsOfPackets(tc.input)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}

func TestCalculateValueOfOutermostPacket(t *testing.T) {
	type TestCase struct {
		input          string
		expectedOutput int
	}
	testCases := []TestCase{
		//{
		//	input:          "C200B40A82",
		//	expectedOutput: 3,
		//},
		//{
		//	input:          "04005AC33890",
		//	expectedOutput: 54,
		//},
		{
			input:          "880086C3E88112",
			expectedOutput: 7,
		},
		{
			input:          "CE00C43D881120",
			expectedOutput: 9,
		},
		//{
		//	input:          "D8005AC2A8F0",
		//	expectedOutput: 1,
		//},
		//{
		//	input:          "F600BC2D8F",
		//	expectedOutput: 0,
		//},
		//{
		//	input:          "9C005AC2F8F0",
		//	expectedOutput: 0,
		//},
		//{
		//	input:          "9C0141080250320F1802104A08",
		//	expectedOutput: 1,
		//},
	}

	for id, tc := range testCases {
		actualResult := CalculateValueOfOutermostPacket(tc.input)
		if actualResult != tc.expectedOutput {
			t.Errorf("[ID %d] Got %v; Want %v", id+1, actualResult, tc.expectedOutput)
		}
	}
}
