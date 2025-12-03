package main

import (
	"flag"
	"log"

	"github.com/hudsn/aoc25/day1"
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
	1: day1.Day1Solve,
}
