package main

import (
	"flag"
	"log"

	"github.com/hudsn/aoc25/day1"
	"github.com/hudsn/aoc25/day2"
	"github.com/hudsn/aoc25/day3"
	"github.com/hudsn/aoc25/day4"
	"github.com/hudsn/aoc25/day5"
	"github.com/hudsn/aoc25/day6"
	"github.com/hudsn/aoc25/day7"
)

func main() {

	day := flag.Int("day", 1, "day to run aoc for")
	flag.Parse()
	if solveFn, ok := dayMap[*day]; ok {
		solveFn()
		return
	}
	log.Fatalf("day currently not supported/solved: %d", *day)
}

var dayMap = map[int]func(){
	1: day1.Solve,
	2: day2.SolveP2,
	3: day3.SolveP2,
	4: day4.Solve,
	5: day5.Solve,
	6: day6.Solve,
	7: day7.Solve,
}
