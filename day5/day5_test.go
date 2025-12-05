package day5

import (
	"fmt"
	"log"
	"testing"
)

func TestP2(t *testing.T) {
	tests := []testCase{
		{
			in: []byte(`3-5
10-14
16-20
12-18`),
			want: 14,
		},
		{
			in: []byte(`1-2
3-4
5-5`),
			want: 5,
		},
	}

	for _, tt := range tests {
		freshRanges := buildRanges(tt.in)
		if freshRanges == nil {
			log.Fatal("issue parsing id ranges")
		}
		ranges := mergeRangesR(freshRanges)
		fmt.Println(ranges)
		got := sumListRangeSlots(ranges)
		if tt.want != got {
			t.Errorf("incorrect outcome: want=%d. got=%d", tt.want, got)
		}
	}

}

type testCase struct {
	in   []byte
	want int
}
