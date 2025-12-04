package day3

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"

	_ "embed"
)

//go:embed day3.txt
var inputBytes []byte

// 0 - 9 => sorted list of idx (low to high)
type batteryIndex struct {
	battery []int
	index   map[int][]int
}

func SolveP1() {
	scanner := bufio.NewScanner(bytes.NewBuffer(inputBytes))
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
		gotIdx, gotNum := batteryIdx.getHighestNumber(-1)
		_, gotNum2 := batteryIdx.getHighestNumber(gotIdx)
		toAdd, err := strconv.Atoi(strconv.Itoa(gotNum) + strconv.Itoa(gotNum2))
		if err != nil {
			log.Fatalf("unable to convert and combine battery numbers to integers: %s", err)
		}
		total += toAdd
		batteryNum++
	}

	fmt.Printf("FLAG: %d\n", total)
}

// return index, value
func (bi batteryIndex) getHighestNumber(afterIndex int) (int, int) {
	for i := 9; i > 0; i-- {
		indexList, ok := bi.index[i]
		if !ok {
			continue
		}
		for _, batteryPosition := range indexList {
			if afterIndex == -1 && batteryPosition == len(bi.battery)-1 { // can't use last index as first item
				continue
			}
			if batteryPosition > afterIndex {
				return batteryPosition, i
			}
		}
	}
	return -1, -1
}

func buildIndex(arr []int) batteryIndex {
	ret := batteryIndex{
		battery: arr,
		index:   make(map[int][]int),
	}
	for idx, v := range arr {
		if _, ok := ret.index[v]; !ok {
			ret.index[v] = []int{}
		}
		ret.index[v] = append(ret.index[v], idx)
	}
	// sort
	for _, v := range ret.index {
		slices.Sort(v)
	}
	return ret
}
