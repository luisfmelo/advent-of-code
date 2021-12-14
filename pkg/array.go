package pkg

import "github.com/luisfmelo/go-advent-of-code-2021/pkg/types"

func CopyArrString(arr []string) []string {
	return append([]string{}, arr...)
}

func MergeMaps(original, mapToMerge map[string]bool) map[string]bool {
	for k, v := range mapToMerge {
		original[k] = v
	}

	return original
}

func CopyMatrix(m types.Matrix) types.Matrix {
	var newMatrix = make(types.Matrix, len(m))
	for i := range m {
		newMatrix[i] = make([]int, len(m[i]))
		for j := range m[i] {
			newMatrix[i][j] = m[i][j]
		}
	}

	return newMatrix
}
