package pkg

import (
	"bufio"
	"bytes"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg/types"
	"io"
	"strconv"
	"strings"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

// ReadLines reads \n separated strings from r. If there's an error, it
// returns the lines successfully read so far as well as the error value.
func ReadLines(scanner *bufio.Scanner) ([]string, error) {
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

// ReadLine reads 1 line until the first \n, it returns the line read as well as the error value
func ReadLine(scanner *bufio.Scanner) (string, error) {
	scanner.Scan()
	return scanner.Text(), scanner.Err()
}

// ReadIntegersInLine reads 1 line until the first \n, it returns the list of integers on that line
// as well as the error value
func ReadIntegersInLine(scanner *bufio.Scanner, sep string) ([]int, error) {
	scanner.Scan()

	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var integerArr []int
	for _, n := range strings.Split(line, sep) {
		integer, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		integerArr = append(integerArr, integer)
	}

	return integerArr, nil
}

func ReadIntsByDelimiter(r io.Reader, delimiter string) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(SplitAt(delimiter))
	var result []int
	for scanner.Scan() {
		result = append(result, StrToInt(scanner.Text()))
	}
	return result, scanner.Err()
}

func SplitAt(substring string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	searchBytes := []byte(substring)
	searchLen := len(searchBytes)
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)

		// Return nothing if at end of file and no data passed
		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}

		// Find next separator and return token
		if i := bytes.Index(data, searchBytes); i >= 0 {
			return i + searchLen, data[0:i], nil
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return dataLen, data, nil
		}

		// Request more data.
		return 0, nil, nil
	}
}

// ReadMatrix reads \n separated strings from r. And then each character.
// If there's an error, it returns the lines successfully read so far as well as the error value.
func ReadMatrix(scanner *bufio.Scanner) ([][]rune, error) {
	var result [][]rune
	for scanner.Scan() {
		var row []rune
		for _, r := range scanner.Text() {
			row = append(row, r)
		}
		result = append(result, row)
	}

	return result, scanner.Err()
}

// ReadIntMatrix reads \n separated strings from r. And then each character.
// If there's an error, it returns the lines successfully read so far as well as the error value.
func ReadIntMatrix(scanner *bufio.Scanner) (types.Matrix, error) {
	var matrix types.Matrix
	rowIndex := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		matrix = append(matrix, []int{})
		for columnIndex, v := range strings.Fields(scanner.Text()) {
			matrix[rowIndex] = append(matrix[rowIndex], 0)
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			matrix[rowIndex][columnIndex] = n
		}
		rowIndex++
	}
	if len(matrix) == 0 {
		return nil, io.EOF
	}

	return matrix, scanner.Err()
}

// ReadDigitMatrixWithoutSeparator reads \n separated strings from r.
// It will identify each digit (0 to 9) and output a matrix of ints.
func ReadDigitMatrixWithoutSeparator(scanner *bufio.Scanner) (types.Matrix, error) {
	var matrix types.Matrix
	rowIndex := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		matrix = append(matrix, []int{})
		for _, v := range strings.Fields(scanner.Text()) {
			for columnIndex, r := range v {
				matrix[rowIndex] = append(matrix[rowIndex], 0)
				n, err := strconv.Atoi(string(r))
				if err != nil {
					return nil, err
				}
				matrix[rowIndex][columnIndex] = n
			}
		}
		rowIndex++
	}
	if len(matrix) == 0 {
		return nil, io.EOF
	}

	return matrix, scanner.Err()
}

// ReadMultipleIntMatrix reads \n separated int matrixes from r.
// If there's an error, it returns the lines successfully read so far as well as the error value.
func ReadMultipleIntMatrix(scanner *bufio.Scanner) ([]types.Matrix, error) {
	var err error
	var matrixes []types.Matrix
	for {
		m, err := ReadIntMatrix(scanner)
		if err != nil {
			break
		}
		matrixes = append(matrixes, m)
	}
	if err != nil && err != io.EOF {
		return nil, err
	}

	return matrixes, nil
}
