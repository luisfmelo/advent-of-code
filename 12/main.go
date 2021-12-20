package main

import (
	"bufio"
	"fmt"
	"github.com/luisfmelo/go-advent-of-code-2021/pkg"
	"log"
	"os"
	"strings"
)

const inputPath = "12/input.txt"

const (
	splitterChar = "-"
)

type passagePath struct {
	path            []string
	smallCavesCache map[string]int
	directionsTaken map[string]bool
}

type passagePathQueue struct {
	passagePaths []passagePath
}

func (q *passagePathQueue) Push(e passagePath) {
	q.passagePaths = append(q.passagePaths, e)
}

func (q *passagePathQueue) Pop() passagePath {
	e := q.passagePaths[0]

	if len(q.passagePaths) == 1 {
		q.passagePaths = []passagePath{}
	} else {
		q.passagePaths = q.passagePaths[1:]
	}

	return e
}

func (q *passagePathQueue) Size() int {
	return len(q.passagePaths)
}

func isSmallCave(cave string) bool {
	return cave == strings.ToLower(cave)
}

func CountNumberOfPaths(inputLines []string) int {
	connections := map[string][]string{}
	for _, inputLine := range inputLines {
		splitter := strings.Split(inputLine, splitterChar)

		if _, exists := connections[splitter[0]]; !exists {
			connections[splitter[0]] = []string{}
		}
		connections[splitter[0]] = append(connections[splitter[0]], splitter[1])

		if _, exists := connections[splitter[1]]; !exists {
			connections[splitter[1]] = []string{}
		}
		connections[splitter[1]] = append(connections[splitter[1]], splitter[0])
	}

	var possiblePaths []string
	queue := passagePathQueue{
		passagePaths: []passagePath{
			{
				path:            []string{"start"},
				smallCavesCache: map[string]int{"start": 1},
				directionsTaken: map[string]bool{},
			},
		},
	}

	for queue.Size() > 0 {
		p := queue.Pop()
		lastCave := p.path[len(p.path)-1]
		for _, connectedCave := range connections[lastCave] {
			smallCavesCache := p.smallCavesCache
			_, beenHere := smallCavesCache[connectedCave]
			if isSmallCave(connectedCave) {
				if beenHere {
					continue
				}
				smallCavesCache[connectedCave] = 1
			}

			directionsTaken := p.directionsTaken
			connectionString := fmt.Sprintf("%s->%s", lastCave, connectedCave)
			if _, alreadyDidThis := directionsTaken[connectionString]; alreadyDidThis {
				continue
			}
			directionsTaken[connectionString] = true

			passagePathElement := passagePath{
				path:            append(p.path, connectedCave),
				smallCavesCache: smallCavesCache,
				directionsTaken: directionsTaken,
			}

			if connectedCave == "end" {
				possiblePaths = append(possiblePaths, strings.Join(passagePathElement.path, "->"))

				continue
			}

			queue.Push(passagePathElement)
		}
	}

	return len(possiblePaths)
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
		func() string { return fmt.Sprintf("%v", CountNumberOfPaths(lines)) },
		func() string { return fmt.Sprintf("%v", CountNumberOfPaths(lines)) },
	)
}
