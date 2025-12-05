package day4

import "testing"

func TestPart1(t *testing.T) {
	cases := []testCase{
		{
			in: []byte(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`),
			want: 13,
		},
	}

	for _, tt := range cases {
		got := solve(tt.in)
		if tt.want != got {
			t.Errorf("wrong outcome. want=%d. got=%d", tt.want, got)
		}
	}
}

type testCase struct {
	in   []byte
	want int
}
