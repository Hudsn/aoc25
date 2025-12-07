package day7

import (
	"bufio"
	"bytes"
	"testing"
)

func TestP1(t *testing.T) {
	cases := []testCase{
		{
			inputBytes: []byte(`.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`),
			want: 21,
		},
	}
	for _, tt := range cases {
		scanner := bufio.NewScanner(bytes.NewBuffer(tt.inputBytes))
		state := rowState{
			currentState:  []nodeType{},
			splitCount:    0,
			timelineCount: []int{},
		}
		for scanner.Scan() {
			state.processLine(scanner.Text())
		}

		if tt.want != state.splitCount {
			t.Errorf("wrong outcome: want=%d. got=%d", tt.want, state.splitCount)
		}
	}

}

type testCase struct {
	inputBytes []byte
	want       int
}
