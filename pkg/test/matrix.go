package test

import "github.com/luisfmelo/go-advent-of-code-2021/pkg/types"

func MatrixIsEqual(m1, m2 types.Matrix) bool {
	if m1.NumberOfRows() != m2.NumberOfRows() {
		return false
	}

	if m1.NumberOfColumns() != m2.NumberOfColumns() {
		return false
	}

	for i := range m1 {
		for j := range m1[i] {
			if m1[i][j] != m2[i][j] {
				return false
			}
		}
	}

	return true
}
