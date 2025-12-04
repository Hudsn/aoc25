package day3

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func solveP2(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	batteryNum := 0
	total := 0
	for scanner.Scan() {
		batteryList := []int{}
		for idx, char := range scanner.Text() {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatalf("unable to convert battery %d entry at idx %d to a number: %s", batteryNum, idx, err)
			}
			batteryList = append(batteryList, num)
		}
		batteryIdx := buildIndex(batteryList)

		sortedIdx := getHighest12(batteryIdx)
		str := batteryIdx.buildFinalString(sortedIdx)
		jolt, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}

		total += jolt
		batteryNum++
	}
	return total
}
func SolveP2() {
	fmt.Printf("FLAG: %d\n", solveP2(inputBytes))
}

func (bi batteryIndex) buildFinalString(sortedIndexList []int) string {
	retStr := []string{}
	for _, idx := range sortedIndexList {
		toAdd := strconv.Itoa(bi.battery[idx])
		retStr = append(retStr, toAdd)
	}
	return strings.Join(retStr, "")
}

// return idx
func (bi batteryIndex) getHighest12Digit(currentIdx []int) int {
	reserveSpaces := 12 - len(currentIdx)
	minIdx := -1
	if len(currentIdx) >= 1 {
		minIdx = currentIdx[len(currentIdx)-1]
	}
	for i := 9; i > 0; i-- {
		indexList, ok := bi.index[i]
		if !ok {
			continue
		}
		for _, batteryPosition := range indexList {
			if batteryPosition > len(bi.battery)-reserveSpaces {
				continue // need to pick a next number that still has enough space for the rest of the sequence
			}
			if slices.Contains(currentIdx, batteryPosition) {
				continue // don't reuse any numbers
			}
			if batteryPosition > minIdx {
				return batteryPosition //only use numbers occurring after our current sequence
			}
		}
	}
	return -1
}

func getHighest12(bi batteryIndex) []int {
	indexList := []int{}
	for len(indexList) < 12 {
		minIdx := bi.getHighest12Digit(indexList)
		indexList = append(indexList, minIdx)
	}
	slices.Sort(indexList)
	return indexList
}
