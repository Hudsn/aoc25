package day5

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed day5.txt
var inputBytes []byte

func solve(input []byte) int {
	rangeBytes, idBytes := splitOriginalInput(input)
	_ = idBytes // part 2
	if rangeBytes == nil {
		log.Fatal("issue splitting original input")
	}
	freshRanges := buildRanges(rangeBytes)
	if freshRanges == nil {
		log.Fatal("issue parsing id ranges")
	}
	//part 1
	// ret, err := countFreshIds(idBytes, freshRanges)
	// if err != nil {
	// 	log.Fatalf("unable to process input ids: %s", err)
	// 	return 0
	// }
	// return ret

	return sumListRangeSlots(mergeRangesR(freshRanges))
}

func Solve() {
	fmt.Printf("FLAG: %d\n", solve(inputBytes))
}

type freshRange struct {
	min int
	max int
}

// part 1
// func countFreshIds(idBytes []byte, freshRanges []freshRange) (int, error) {
// 	count := 0
// 	scanner := bufio.NewScanner(bytes.NewBuffer(idBytes))
// 	for scanner.Scan() {
// 		num, err := strconv.Atoi(scanner.Text())
// 		if err != nil {
// 			return 0, fmt.Errorf("issue parsing ")
// 		}
// 		for _, idRange := range freshRanges {
// 			if idRange.min < num && num < idRange.max {
// 				count++
// 				break
// 			}
// 		}
// 	}
// 	return count, nil
// }

func buildRanges(rangeBytes []byte) []freshRange {
	ret := []freshRange{}
	scanner := bufio.NewScanner(bytes.NewBuffer(rangeBytes))
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		split := strings.Split(text, "-")
		if len(split) != 2 {
			return nil
		}
		min, err := strconv.Atoi(split[0])
		if err != nil {
			return nil
		}
		max, err := strconv.Atoi(split[1])
		if err != nil {
			return nil
		}
		ret = append(ret, freshRange{min: min, max: max})
	}
	return ret
}

func splitOriginalInput(input []byte) ([]byte, []byte) {
	split := bytes.Split(input, []byte("\n\n"))
	if len(split) != 2 {
		return nil, nil
	}
	return split[0], split[1]
}

func sumListRangeSlots(ranges []freshRange) int {
	sum := 0
	for _, r := range ranges {
		sum += (r.max - r.min + 1) // +1 because of inclusive ids; ex: 1 to 3 would be 3
	}
	return sum
}

func mergeRangesR(ranges []freshRange) []freshRange {
	prevLen := len(ranges)
	ret := ranges
	for {
		ret = mergeRanges(ret)
		if len(ret) == prevLen {
			break
		}
		prevLen = len(ret)
	}
	return ret
}

func mergeRanges(freshRanges []freshRange) []freshRange {
	newRange := []freshRange{}
	skipList := []int{}
	for i := range freshRanges {
		if slices.Contains(skipList, i) {
			continue
		}
		r1 := freshRanges[i]
		toAdd := r1
		for j := range freshRanges {
			if i == j {
				continue
			}
			r2 := freshRanges[j]
			if isRangeOverlap(r1, r2) {
				toAdd = consolidateRanges(r1, r2)
				skipList = append(skipList, j)
				break
			}
		}
		newRange = append(newRange, toAdd)
	}
	return newRange
}

func isRangeOverlap(r1 freshRange, r2 freshRange) bool {
	return r1.min <= r2.max && r2.min <= r1.max
}

func consolidateRanges(r1 freshRange, r2 freshRange) freshRange {
	var min int
	var max int
	if r1.min < r2.min {
		min = r1.min
	} else {
		min = r2.min
	}
	if r1.max > r2.max {
		max = r1.max
	} else {
		max = r2.max
	}
	return freshRange{min: min, max: max}
}
