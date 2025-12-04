package day3

import "testing"

func TestP2(t *testing.T) {
	tests := []testCase{
		{
			in:   "987654321111111",
			want: 987654321111,
		},
		{
			in:   "811111111111119",
			want: 811111111119,
		},
		{
			in:   "811111111111119\n987654321111111",
			want: 811111111119 + 987654321111,
		},
		{
			in:   "818181911112111",
			want: 888911112111,
		},
		{
			in:   "234234234234278",
			want: 434234234278,
		},
	}
	for _, tt := range tests {
		got := solveP2([]byte(tt.in))
		if tt.want != got {
			t.Errorf("wrong outcome. want=%d. got=%d", tt.want, got)
		}
	}

}

type testCase struct {
	in   string
	want int
}
