package pkg

import "github.com/luisfmelo/go-advent-of-code-2021/pkg/types"

func IntArrIsEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func MatrixElementArrIsEqual(arr1, arr2 []types.MatrixElement) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i].ToString() != arr2[i].ToString() {
			return false
		}
	}

	return true
}
