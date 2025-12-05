package day4

import (
	"bufio"
	"bytes"
	"fmt"

	_ "embed"
)

//go:embed day4.txt
var inputBytes []byte

func solve(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	grid := newGrid()
	for scanner.Scan() {
		grid.newRow(scanner.Text())
	}
	// part_1
	// return countSpacesWithFewerThanXAdjancent(grid, 4)
	return removeCoords(grid, 4)
}

func Solve() {
	fmt.Printf("FLAG: %d\n", solve(inputBytes))
}

func removeCoords(g grid, targetNum int) int {
	ret := 0
	removed := -1
	for {
		removed = countRemovals(g, targetNum)
		if removed == 0 || removed == -1 {
			break
		}
		ret += removed
	}
	return ret
}

func countRemovals(g grid, targetNum int) int {
	removedCount := 0
	for _, row := range g.rows {
		for _, coord := range row {
			if coord.isFull {
				surrounding := g.countSurrounding(coord)
				if surrounding < targetNum {
					g.replaceCoord(coordinate{
						row:    coord.row,
						col:    coord.col,
						isFull: false,
					})
					removedCount++
				}
			}
		}
	}
	return removedCount
}

//part 1
// func countSpacesWithFewerThanXAdjancent(g grid, targetNum int) int {
// 	ret := 0
// 	for _, row := range g.rows {
// 		for _, coord := range row {
// 			if coord.isFull {
// 				surrounding := g.countSurrounding(coord)
// 				if surrounding < targetNum {
// 					ret++
// 				}
// 			}
// 		}
// 	}
// 	return ret
// }

type coordinate struct {
	col    int
	row    int
	isFull bool
}

type grid struct {
	lastRowNum int // first row = 0, etc...
	rows       [][]coordinate
}

func newGrid() grid {
	return grid{
		lastRowNum: -1,
		rows:       [][]coordinate{},
	}
}

func (g grid) maxRow() int {
	return len(g.rows) - 1
}

func (g grid) maxCol() int {
	return len(g.rows[0]) - 1
}

func (g *grid) newRow(rowText string) {
	rowNum := g.lastRowNum + 1
	newRow := []coordinate{}
	for idx, char := range rowText {
		toAdd := coordinate{
			col:    idx,
			row:    rowNum,
			isFull: char == '@',
		}
		newRow = append(newRow, toAdd)
	}
	g.rows = append(g.rows, newRow)
	g.lastRowNum = rowNum
}

func (g grid) getCoordinate(row int, col int) (coordinate, bool) {
	if row < 0 || col < 0 {
		return coordinate{}, false
	}
	if row > g.maxRow() || col > g.maxCol() {
		return coordinate{}, false
	}
	return g.rows[row][col], true
}

func (g grid) countSurrounding(c coordinate) int {
	ret := 0
	for _, coord := range g.getValidSurroundingCoords(c) {
		if coord.isFull {
			ret++
		}
	}
	return ret
}

func (g grid) getValidSurroundingCoords(c coordinate) []coordinate {
	ret := []coordinate{}
	for ri := c.row - 1; ri <= c.row+1; ri++ {
		for ci := c.col - 1; ci <= c.col+1; ci++ {
			if ri == c.row && ci == c.col {
				continue // dont count self
			}
			if coord, isValid := g.getCoordinate(ri, ci); isValid {
				ret = append(ret, coord)
			}
		}
	}
	return ret
}

// part2
func (g *grid) replaceCoord(c coordinate) {
	g.rows[c.row][c.col] = c
}
