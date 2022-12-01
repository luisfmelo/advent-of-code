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

const inputPath = "2021/10/input.txt"

const (
	roundBracketErrorScore  = 3
	squareBracketErrorScore = 57
	curlyBracketErrorScore  = 1197
	angleBracketErrorScore  = 25137

	roundBracketCompletionScore  = 1
	squareBracketCompletionScore = 2
	curlyBracketCompletionScore  = 3
	angleBracketCompletionScore  = 4
)

type score struct {
	errorScore      int
	completionScore int
}

type bracket rune

const (
	roundBracketOpener  bracket = '('
	roundBracketCloser  bracket = ')'
	squareBracketOpener bracket = '['
	squareBracketCloser bracket = ']'
	curlyBracketOpener  bracket = '{'
	curlyBracketCloser  bracket = '}'
	angleBracketOpener  bracket = '<'
	angleBracketCloser  bracket = '>'
)

var (
	openerBrackets = map[bracket]bool{
		roundBracketOpener:  true,
		squareBracketOpener: true,
		curlyBracketOpener:  true,
		angleBracketOpener:  true,
	}
	closerBrackets = map[bracket]bool{
		roundBracketCloser:  true,
		squareBracketCloser: true,
		curlyBracketCloser:  true,
		angleBracketCloser:  true,
	}
)

func isOpenerBracket(b bracket) bool {
	_, exists := openerBrackets[b]
	return exists
}

func isCloserBracket(b bracket) bool {
	_, exists := closerBrackets[b]
	return exists
}

func closerBracketMatchesOpenerBracket(opener, closer bracket) bool {
	switch opener {
	case roundBracketOpener:
		return closer == roundBracketCloser
	case squareBracketOpener:
		return closer == squareBracketCloser
	case curlyBracketOpener:
		return closer == curlyBracketCloser
	case angleBracketOpener:
		return closer == angleBracketCloser
	}

	return false
}

func calculateLineScore(line string) score {
	stack := types.NewRuneStack()
	for _, r := range line {
		b := bracket(r)
		if isOpenerBracket(b) {
			stack.Push(r)
		} else if isCloserBracket(b) {
			lastOpened, err := stack.Pop()
			if err != nil {
				switch err.(type) {
				case types.ErrEmptyStack:
					panic(err.Error())
				default:
					panic(err.Error())
				}
			}

			if !closerBracketMatchesOpenerBracket(bracket(lastOpened), b) {
				switch b {
				case roundBracketCloser:
					return score{errorScore: roundBracketErrorScore}
				case squareBracketCloser:
					return score{errorScore: squareBracketErrorScore}
				case curlyBracketCloser:
					return score{errorScore: curlyBracketErrorScore}
				case angleBracketCloser:
					return score{errorScore: angleBracketErrorScore}
				}
			}
		}
	}

	completionScore := 0
	for stack.Size() > 0 {
		r, err := stack.Pop()
		if err != nil {
			panic(err)
		}

		switch bracket(r) {
		case roundBracketOpener:
			completionScore = completionScore*5 + roundBracketCompletionScore
		case squareBracketOpener:
			completionScore = completionScore*5 + squareBracketCompletionScore
		case curlyBracketOpener:
			completionScore = completionScore*5 + curlyBracketCompletionScore
		case angleBracketOpener:
			completionScore = completionScore*5 + angleBracketCompletionScore
		}
	}

	return score{completionScore: completionScore}
}

func CalculateSyntaxErrorScore(lines []string) int {
	var errorScore int
	for _, line := range lines {
		errorScore += calculateLineScore(line).errorScore
	}

	return errorScore
}

func GetMiddleCompletionScore(lines []string) int {
	var completionScores []int
	for _, line := range lines {
		completionScore := calculateLineScore(line).completionScore
		if completionScore > 0 {
			completionScores = append(completionScores, completionScore)
		}
	}

	sort.Ints(completionScores)

	return completionScores[len(completionScores)/2]
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

	lines, err := pkg.ReadLines(scanner)
	pkg.PanicErr(err)

	pkg.RunWithTime(
		func() string { return fmt.Sprintf("%v", CalculateSyntaxErrorScore(lines)) },
		func() string { return fmt.Sprintf("%v", GetMiddleCompletionScore(lines)) },
	)
}
