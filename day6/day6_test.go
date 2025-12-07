package day6

import (
	"reflect"
	"testing"
)

type testCaseFull struct {
	originalInput []byte
	want          int
}

func TestPart2Full(t *testing.T) {
	cases := []testCaseFull{
		{

			// The rightmost problem is 4 + 431 + 623 = 1058
			// The second problem from the right is 175 * 581 * 32 = 3253600
			// The third problem from the right is 8 + 248 + 369 = 625
			// Finally, the leftmost problem is 356 * 24 * 1 = 8544

			originalInput: []byte(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `),
			want: 3263827,
		},
	}
	for _, tt := range cases {
		got := solve(tt.originalInput)
		if tt.want != got {
			t.Errorf("wrong outcome: want=%d got=%d", tt.want, got)
		}
	}
}

type mathCases struct {
	originalInput []byte
	want          []int
}

func TestMath(t *testing.T) {
	cases := []mathCases{
		{

			// The rightmost problem is 4 + 431 + 623 = 1058
			// The second problem from the right is 175 * 581 * 32 = 3253600
			// The third problem from the right is 8 + 248 + 369 = 625
			// Finally, the leftmost problem is 356 * 24 * 1 = 8544
			originalInput: []byte(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `),
			want: []int{
				8544, 625, 3253600, 1058,
			},
		},
	}
	for _, tt := range cases {
		p := buildColBasedProblems(tt.originalInput)
		for idx, want := range tt.want {
			prob := p[idx]
			got := prob.eval()
			if want != prob.eval() {
				t.Errorf("wrong outcome: want=%d got=%d", tt.want, got)
			}

		}
	}
}

func TestWeirdNums(t *testing.T) {

	// 	123 328  51 64
	//  45 64  387 23
	//   6 98  215 314
	// *   +   *   +
	cases := []testCaseWeirdNums{
		{

			// The rightmost problem is 4 + 431 + 623 = 1058
			// The second problem from the right is 175 * 581 * 32 = 3253600
			// The third problem from the right is 8 + 248 + 369 = 625
			// Finally, the leftmost problem is 356 * 24 * 1 = 8544

			originalInput: []byte(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `),
			want: [][]int{
				{356, 24, 1},
				{8, 248, 369},
				{175, 581, 32},
				{4, 431, 623},
			},
		},
	}
	for _, tt := range cases {
		p := buildColBasedProblems(tt.originalInput)
		for idx, want := range tt.want {
			prob := p[idx]
			if !reflect.DeepEqual(prob.nums, want) {
				t.Errorf("wrong outcome.\nwant:\n\t%#v\ngot:\n\t%#v\n", want, prob.nums)
			}
		}

	}

}

type testCaseWeirdNums struct {
	originalInput []byte
	want          [][]int
}
