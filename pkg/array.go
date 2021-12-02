package pkg

func CopyArrString(arr []string) []string {
	return append([]string{}, arr...)
}

func MergeMaps(original, mapToMerge map[string]bool) map[string]bool {
	for k, v := range mapToMerge {
		original[k] = v
	}
	return original
}
