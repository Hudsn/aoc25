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
	left direction = iota
	right
)

type instruction struct {
	direction direction
	count     int
}

type dial struct {
	pointer int
}

func (d *dial) countZeroClicks(ins instruction) int {
	distanceToZero := 100 - d.pointer
	if ins.direction == left {
		distanceToZero = d.pointer
		if distanceToZero == 0 {
			distanceToZero = 100
		}
	}
	diff := distanceToZero - ins.count
	if diff <= 0 {
		diff *= -1
		return 1 + (diff-(diff%100))/100
	}
	return 0
}

func (d *dial) updatePointer(ins instruction) {
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

func solve(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	counter := 0
	dial := dial{
		pointer: 50,
	}
	for scanner.Scan() {
		instruction := textToInstruction(scanner.Text())
		counter += dial.countZeroClicks(instruction)
		dial.updatePointer(instruction)
	}
	return counter
}

func Solve() {
	fmt.Printf("FLAG: %d\n", solve(day1Bytes))
}
