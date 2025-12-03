package day1

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"

	_ "embed"
)

//go:embed day1.txt
var day1Bytes []byte

type direction int

const (
	left = iota
	right
)

type instruction struct {
	direction direction
	count     int
}

type dial struct {
	pointer int
}

func (d *dial) processInstruction(ins instruction) {
	if ins.direction == left {
		d.moveRight(100 - (ins.count % 100)) // on a 100 position dial, moving right 25 == moving left 75
	} else {
		d.moveRight(ins.count)
	}
}

func (d *dial) moveRight(count int) {
	d.pointer = (d.pointer + count) % 100
}

func textToInstruction(text string) instruction {
	directionChar := string(text[0])
	var direction direction
	switch directionChar {
	case "R":
		direction = right
	case "L":
		direction = left
	}

	count, err := strconv.Atoi(string(text[1:]))
	if err != nil {
		log.Fatalf("instruction conversion error: %s", err)
	}

	return instruction{
		direction: direction,
		count:     count,
	}
}

func Day1Solve() {
	scanner := bufio.NewScanner(bytes.NewBuffer(day1Bytes))
	counter := 0
	dial := dial{
		pointer: 50,
	}
	for scanner.Scan() {
		instruction := textToInstruction(scanner.Text())
		dial.processInstruction(instruction)
		if dial.pointer == 0 {
			counter++
		}
	}
	fmt.Printf("FLAG: %d\n", counter)
}
