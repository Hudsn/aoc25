package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func SolveP2() {
	fmt.Printf("FLAG: %d\n", solveP2(string(inputBytes)))
}

func solveP2(input string) int {
	count := 0
	rangeStrings := strings.Split(input, ",")
	for _, idRangeStr := range rangeStrings {
		idRangeStr = strings.TrimSpace(idRangeStr)
		start, end := stringRangeToStartEnd(idRangeStr)
		if start == -1 && end == -1 {
			log.Fatalf("unable to convert Id entry %q to start and end range", idRangeStr)
			return -1
		}
		count += sumWithBruteForce(start, end)
	}
	return count
}

func stringRangeToStartEnd(rangeStr string) (int, int) {
	splits := strings.Split(rangeStr, "-")
	if len(splits) != 2 {
		return -1, -1
	}
	lowStr := splits[0]
	highStr := splits[1]
	min, err := strconv.Atoi(lowStr)
	if err != nil {
		return -1, -1
	}
	max, err := strconv.Atoi(highStr)
	if err != nil {
		return -1, -1
	}
	return min, max
}

func sumWithBruteForce(start int, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if checkRepeatingNumber(i) {
			count += i
		}
	}
	return count
}

func checkRepeatingNumber(num int) bool {
	asStr := strconv.Itoa(num)
	if len(asStr) == 1 {
		return false
	}
	var end int
	if len(asStr)%2 == 0 {
		end = len(asStr) / 2
	} else {
		end = (len(asStr) - 1) / 2
	}
	for i := 1; i <= end; i++ {
		if len(asStr)%i != 0 { //skip any lengths that can't perfectly repeat
			continue
		}
		repCount := len(asStr) / i
		chunk := asStr[:i]
		wantMatch := strings.Repeat(chunk, repCount)
		if wantMatch == asStr {
			return true
		}
	}
	return false
}
