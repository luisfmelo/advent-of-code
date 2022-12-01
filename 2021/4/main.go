package main

import (
	"bufio"
	"fmt"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
	"log"
	"os"

	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
)

const (
	cross     = -1
	inputPath = "2021/4/input.txt"
)

func checkRowWin(board types.Matrix, rowIdx int) bool {
	for _, v := range board[rowIdx] {
		if v != cross {
			return false
		}
	}

	return true
}

func checkColumnWin(board types.Matrix, columnIdx int) bool {
	for _, row := range board {
		if row[columnIdx] != cross {
			return false
		}
	}

	return true
}

func markNumber(board types.Matrix, numberDrew int) (types.Matrix, bool) {
	for i, row := range board {
		for j, col := range row {
			if col == numberDrew {
				board[i][j] = cross
				return board, checkRowWin(board, i) || checkColumnWin(board, j)
			}
		}
	}
	return board, false
}

func calculateScoreOfBoard(board types.Matrix, lastNumberDrew int) int {
	sumNumbersUnmarked := 0
	for _, row := range board {
		for _, col := range row {
			if col == cross {
				continue
			}
			sumNumbersUnmarked += col
		}
	}

	return sumNumbersUnmarked * lastNumberDrew
}

func CalculateScoreOfTheWinningBoard(numbersDrew []int, boards []types.Matrix) int {
	var (
		winnerBoard    *types.Matrix
		win            bool
		lastNumberDrew int
	)

	for _, lastNumberDrew = range numbersDrew {
		// mark number in all boards & check for winning boards
		for i, board := range boards {
			if boards[i], win = markNumber(board, lastNumberDrew); win {
				winnerBoard = &boards[i]
			}
		}

		if winnerBoard != nil {
			break
		}
	}

	return calculateScoreOfBoard(*winnerBoard, lastNumberDrew)
}

func removeFromArr(boards []types.Matrix, mapIndexes map[int]bool) []types.Matrix {
	var remaining []types.Matrix
	for i, board := range boards {
		if _, toDelete := mapIndexes[i]; toDelete {
			continue
		}
		remaining = append(remaining, board)
	}

	return remaining
}

func CalculateScoreOfTheLatestToWinBoard(numbersDrew []int, boards []types.Matrix) int {
	var (
		lastNumberDrew   int
		latestBoardToWin *types.Matrix
	)

	var indexesToDelete = map[int]bool{}
	for _, lastNumberDrew = range numbersDrew {
		indexesToDelete = map[int]bool{}
		// mark number in all boards & check for winning boards
		for i, board := range boards {
			b, win := markNumber(board, lastNumberDrew)
			boards[i] = b
			if win {
				latestBoardToWin = &b
				indexesToDelete[i] = true
			}
		}
		boards = removeFromArr(boards, indexesToDelete)
		if len(boards) == 0 {
			break
		}
	}

	return calculateScoreOfBoard(*latestBoardToWin, lastNumberDrew)
}

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Printf("Error occurred: %v", err)
		}
	}()

	file, err := os.Open(inputPath)
	pkg.PanicErr(err)

	r := bufio.NewReader(file)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	numbersDrew, err := pkg.ReadIntegersInLine(scanner, ",")
	pkg.PanicErr(err)

	_, err = pkg.ReadLine(scanner)
	pkg.PanicErr(err)

	boards, err := pkg.ReadMultipleIntMatrix(scanner)
	pkg.PanicErr(err)

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", CalculateScoreOfTheWinningBoard(numbersDrew, boards)) },
		func() string { return fmt.Sprintf("%v", CalculateScoreOfTheLatestToWinBoard(numbersDrew, boards)) },
	)
}
