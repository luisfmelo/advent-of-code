package pkg

import (
	"log"
	"time"
)

type adventDayFunc func() string

func RunWithTime(part1, part2 adventDayFunc) {
	run := func(partNumber int, f adventDayFunc) {
		if f == nil {
			log.Printf("Part #%d is not implemented yet", partNumber)
			return
		}

		start := time.Now()
		log.Printf("Part #%d result: %v\n", partNumber, f())
		log.Printf("1st part took: %s", time.Since(start))
	}

	run(1, part1)
	run(2, part2)
}
