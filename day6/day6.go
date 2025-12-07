package day6

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed day6.txt
var inputBytes []byte

type operation int

const (
	_ operation = iota
	add
	subtract
	multiply
	divide
)

func solve(input []byte) int {
	//part 1
	// probs := newProblemSet(input)
	probs := buildColBasedProblems(input)
	sum := 0
	for _, prob := range probs {
		sum += prob.eval()
	}
	return sum
}

func Solve() {
	fmt.Printf("FLAG: %d\n", solve(inputBytes))
}

type problem struct {
	nums []int
	operation
}

var opMap = map[operation]func(a, b int) int{
	add:      func(a, b int) int { return a + b },
	subtract: func(a, b int) int { return a - b },
	multiply: func(a, b int) int { return a * b },
	divide:   func(a, b int) int { return a / b },
}

func (p problem) eval() int {
	fn := opMap[p.operation]
	result := p.nums[0]
	for _, next := range p.nums[1:] {
		result = fn(result, next)
	}
	return result
}

//part1
// func newProblemSet(input []byte) []*problem {
// 	ret := []*problem{}
// 	scanner := bufio.NewScanner(bytes.NewBuffer(input))
// 	i := 0
// 	for scanner.Scan() {
// 		strings.FieldsFunc(scanner.Text(), func(r rune) bool { return r == ' ' })
// 		fields := strings.Fields(scanner.Text())
// 		if i == 0 {
// 			for range len(fields) {
// 				ret = append(ret, &problem{nums: []int{}})
// 			}
// 		}
// 		if nums, ok := convertToNums(fields); ok {
// 			for i, num := range nums {
// 				prob := ret[i]
// 				prob.nums = append(prob.nums, num)
// 			}
// 		} else if ops, ok := convertToOps(fields); ok {
// 			for i, op := range ops {
// 				prob := ret[i]
// 				prob.operation = op
// 			}
// 		} else {
// 			log.Fatalf("issue reading data into problems: %v", fields)
// 		}
// 		i++
// 	}
// 	return ret
// }

func convertToNums(fields []string) ([]int, bool) {
	ret := []int{}
	for _, s := range fields {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, false
		}
		ret = append(ret, num)
	}
	return ret, true
}

func convertToOps(fields []string) ([]operation, bool) {
	ret := []operation{}
	for _, s := range fields {
		switch s {
		case "+":
			ret = append(ret, add)
		case "-":
			ret = append(ret, subtract)
		case "*":
			ret = append(ret, multiply)
		case "/":
			ret = append(ret, divide)
		default:
			return nil, false
		}
	}
	return ret, true
}

func buildColBasedProblems(inputBytes []byte) []*problem {

	rawcols := []string{}
	scanner := bufio.NewScanner(bytes.NewBuffer(inputBytes))
	allOps := []operation{}

	for scanner.Scan() {
		if strings.ContainsAny(scanner.Text(), "+-/*") {
			allOps, _ = convertToOps(strings.Fields(scanner.Text()))
			break
		}
		split := strings.Split(scanner.Text(), "")
		for idx, char := range split {
			if len(rawcols) <= idx {
				rawcols = append(rawcols, "")
			}
			rawcols[idx] += char
		}
	}
	probs := []*problem{}
	currentProb := &problem{nums: []int{}}
	for _, colStr := range rawcols {
		if len(strings.ReplaceAll(colStr, " ", "")) == 0 {
			slices.Reverse(currentProb.nums)
			probs = append(probs, currentProb)
			currentProb = &problem{nums: []int{}}
			continue
		}
		asInt, err := strconv.Atoi(strings.TrimSpace(colStr))
		if err != nil {
			log.Fatalf("issue generating columns. strconv:  %s", err)
			return nil
		}
		currentProb.nums = append(currentProb.nums, asInt)
	}
	// make sure we add last in-progress problem
	if len(currentProb.nums) != 0 {
		slices.Reverse(currentProb.nums)
		probs = append(probs, currentProb)
	}

	for idx, entry := range probs {
		entry.operation = allOps[idx]
	}

	return probs
}
