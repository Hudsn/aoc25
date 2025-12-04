package day2

import "testing"

func TestPart1(t *testing.T) {
	cases := []testCase{
		{ //both equal
			in:   "11-22",
			want: 11 + 22,
		},
		{ // lower is odd digits, higher is even digits
			in:   "1-13",
			want: 11,
		},
		{ // lower is even digits, higher is odd digits
			in:   "80-100",
			want: 88 + 99,
		},
	}
	for _, c := range cases {
		got := solveP1(c.in)
		if c.want != got {
			t.Errorf("wrong outcome. Want=%d Got=%d", c.want, got)
		}
	}
}
func TestPart2(t *testing.T) {
	cases := []testCase{
		{
			in:   "11-22",
			want: 11 + 22,
		},
		{
			in:   "1-13",
			want: 11,
		},
		{
			in:   "80-111",
			want: 88 + 99 + 111,
		},
		{
			in:   "824824821-824824827",
			want: 824824824,
		},
	}
	for _, c := range cases {
		got := solveP2(c.in)
		if c.want != got {
			t.Errorf("wrong outcome. Want=%d Got=%d", c.want, got)
		}
	}
}

type testCase struct {
	in   string
	want int
}
