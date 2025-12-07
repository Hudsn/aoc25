package day7

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day7.txt
var inputBytes []byte

func solve(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	state := rowState{
		currentState:  []nodeType{},
		splitCount:    0,
		timelineCount: []int{},
	}
	for scanner.Scan() {
		state.processLine(scanner.Text())
	}

	count := 0
	for _, v := range state.timelineCount {
		count += v
	}
	return count
}

func Solve() {
	fmt.Printf("FLAG: %d\n", solve(inputBytes))
}

type nodeType int

const (
	_ nodeType = iota
	beam
	empty
)

type rowState struct {
	rowSize       int
	currentState  []nodeType
	timelineCount []int
	splitCount    int
}

func (rs *rowState) processLine(line string) {
	if strings.Contains(line, "S") {
		rs.rowSize = len(line)
		for range rs.rowSize {
			rs.currentState = append(rs.currentState, empty)
			rs.timelineCount = append(rs.timelineCount, 0)
		}
	}
	split := strings.Split(line, "")
	newRow := []nodeType{}
	newCounts := []int{}
	for range rs.rowSize {
		newRow = append(newRow, empty)
		newCounts = append(newCounts, 0)
	}
	for idx, char := range split {
		// if newRow[idx] == beam {
		// 	continue
		// }
		switch char {
		case "S":
			newRow[idx] = beam
			newCounts[idx] = 1
		case "^":
			newRow[idx] = empty
			if rs.currentState[idx] == beam {
				if idx > 0 {
					newRow[idx-1] = beam
				}
				if idx+1 < len(newRow) {
					newRow[idx+1] = beam
				}
				rs.splitCount++
			}
			if rs.timelineCount[idx] > 0 {
				if idx > 0 {
					newCounts[idx-1] += rs.timelineCount[idx]
				}
				if idx+1 < rs.rowSize {
					newCounts[idx+1] += rs.timelineCount[idx]
				}
			}
		default:
			if newRow[idx] != beam {
				newRow[idx] = rs.currentState[idx]
			}
			if rs.timelineCount[idx] > 0 {
				newCounts[idx] += rs.timelineCount[idx]
			}
		}
	}
	rs.currentState = newRow
	rs.timelineCount = newCounts
}
