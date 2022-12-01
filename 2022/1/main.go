package main

import (
	"bufio"
	"fmt"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
	"log"
	"os"
	"sort"
)

const inputPath = "2022/1/input.txt"

func getTotalCaloriesPerElf(input types.Matrix) []int {
	caloriesPerElf := make([]int, len(input))
	for elfN, elfCalorieBag := range input {
		for _, calorie := range elfCalorieBag {
			caloriesPerElf[elfN] += calorie
		}
	}

	return caloriesPerElf
}

func MaxFoodCarriedBySingleElf(input types.Matrix) int {
	var maxCaloriesCarrying int
	for _, elfCalorieBag := range input {
		caloriesCarrying := 0
		for _, calorie := range elfCalorieBag {
			caloriesCarrying += calorie
		}
		if caloriesCarrying > maxCaloriesCarrying {
			maxCaloriesCarrying = caloriesCarrying
		}
	}

	return maxCaloriesCarrying
}

func MaxFoodCarriedByTop3Elfs(input types.Matrix) int {
	caloriesPerElf := getTotalCaloriesPerElf(input)
	sort.Slice(caloriesPerElf, func(i, j int) bool {
		return caloriesPerElf[i] > caloriesPerElf[j]
	})

	return caloriesPerElf[0] + caloriesPerElf[1] + caloriesPerElf[2]
}

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error occurred: %v", err)
		}
	}()

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	input, err := pkg.ReadMultipleIntSeparatedByBlankLine(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", MaxFoodCarriedBySingleElf(input)) },
		func() string { return fmt.Sprintf("%v", MaxFoodCarriedByTop3Elfs(input)) },
	)
}
