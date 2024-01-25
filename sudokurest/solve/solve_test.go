package solve

import (
	"reflect"
	"testing"

	"github.com/LaoZhuBaba/sudoku_rest/autogen"
)

func Test_isIntInSlice(t *testing.T) {
	type args struct {
		i int
		s []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2_in_8345",
			args: args{
				i: 2,
				s: []int{8, 3, 4, 5},
			},
			want: false,
		},
		{
			name: "2_in_8325",
			args: args{
				i: 2,
				s: []int{8, 3, 2, 5},
			},
			want: true,
		},
		{
			name: "2_in_empty",
			args: args{
				i: 2,
				s: []int{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIntInSlice(tt.args.i, tt.args.s); got != tt.want {
				t.Errorf("isIntInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCandidates(t *testing.T) {
	type args struct {
		s                        []int
		related_elements_indexes []int
		maxValue                 int
	}
	tests := []struct {
		name           string
		args           args
		wantCandidates []int
	}{
		{
			name: "not_0_1_8",
			args: args{
				s:                        []int{1, 2, 3, 0, 5, 6, 7, 8, 9},
				related_elements_indexes: []int{0, 1, 8},
				maxValue:                 9,
			},
			wantCandidates: []int{3, 4, 5, 6, 7, 8},
		},
		{
			name: "no_candidates",
			args: args{
				s:                        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				related_elements_indexes: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			wantCandidates: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCandidates := findCandidates(tt.args.s, tt.args.related_elements_indexes, tt.args.maxValue); !reflect.DeepEqual(gotCandidates, tt.wantCandidates) {
				t.Errorf("findCandidates() = %#v, want %#v", gotCandidates, tt.wantCandidates)
			}
		})
	}
}

func Test_recursiveSolve(t *testing.T) {
	type args struct {
		i                int
		s                []int
		maxIndex         int
		maxValue         int
		related_elements [][]int
	}
	tests := []struct {
		name           string
		args           args
		want           bool
		wantSideEffect []int
	}{
		{
			name: "one_missing_in_middle",
			args: args{
				i:        0,
				s:        []int{1, 2, 3, 0, 5, 6, 7, 8, 9},
				maxIndex: 8,
				maxValue: 9,
				related_elements: [][]int{
					{}, {}, {}, {0, 1, 2, 4, 5, 6, 7, 8}, {}, {}, {}, {}, {},
				},
			},
			want:           true,
			wantSideEffect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "one_missing_in_at_the_end",
			args: args{
				i:        0,
				s:        []int{1, 2, 3, 4, 5, 6, 7, 8, 0},
				maxIndex: 8,
				maxValue: 9,
				related_elements: [][]int{
					{}, {}, {}, {}, {}, {}, {}, {}, {0, 1, 2, 3, 4, 5, 6, 7, 8},
				},
			},
			want:           true,
			wantSideEffect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "one_missing_in_at_the_start",
			args: args{
				i:        0,
				s:        []int{0, 2, 3, 4, 5, 6, 7, 8, 9},
				maxIndex: 8,
				maxValue: 9,
				related_elements: [][]int{
					{1, 2, 3, 4, 5, 6, 7, 8}, {}, {}, {}, {}, {}, {}, {}, {},
				},
			},
			want:           true,
			wantSideEffect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "three_missing",
			args: args{
				i:        0,
				s:        []int{0, 2, 3, 0, 5, 6, 0, 8, 9},
				maxIndex: 8,
				maxValue: 9,
				related_elements: [][]int{
					{1, 2, 4, 5, 7, 8}, {}, {}, {0, 1, 2, 4, 5, 7, 8}, {}, {}, {0, 1, 2, 3, 4, 5, 7, 8}, {}, {},
				},
			},
			want:           true,
			wantSideEffect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "should fail",
			args: args{
				i:        0,
				s:        []int{1, 2, 3, 0, 5, 6, 7, 8, 9},
				maxIndex: 8,
				maxValue: 9,
				related_elements: [][]int{
					{}, {}, {}, {0, 1, 2, 3, 4, 5, 6, 7, 8}, {}, {}, {}, {}, {},
				},
			},
			want:           true,
			wantSideEffect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			// Define a 2 x 2 sudoku with 1 in top left corner.  Each column and each row must contain
			// two unique values max a max value of 2, so the only solution is
			//   1 2
			//   2 1
			name: "solve 2x2",
			args: args{
				i:        0,
				s:        []int{1, 0, 0, 0},
				maxIndex: 3,
				maxValue: 2,
				related_elements: [][]int{
					{1, 2}, {0, 3}, {0, 3}, {1, 2},
				},
			},
			want:           true,
			wantSideEffect: []int{1, 2, 2, 1},
		},
		{
			name: "solve 3x3",
			// 2 0 0        2 3 1
			// 0 2 0   ->   1 2 3
			// 0 1 0        3 1 2
			args: args{
				i:        0,
				s:        []int{2, 0, 0, 0, 2, 0, 0, 1, 0},
				maxIndex: 8,
				maxValue: 3,
				related_elements: [][]int{
					{1, 2, 3, 6}, {0, 2, 4, 7}, {0, 1, 5, 8}, {0, 4, 5, 6}, {1, 3, 5, 7}, {2, 3, 4, 8}, {0, 3, 7, 8}, {1, 4, 6, 8}, {2, 5, 6, 7},
				},
			},
			want:           true,
			wantSideEffect: []int{2, 3, 1, 1, 2, 3, 3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursiveSolve(tt.args.i, tt.args.s, tt.args.maxValue, tt.args.maxIndex, tt.args.related_elements); got != tt.want {
				t.Errorf("recursiveSolve() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.s, tt.wantSideEffect) {
				t.Errorf("recursiveSolve() = %v, wantSideEffect %v", tt.args.s, tt.wantSideEffect)
			}
		})
	}
}

func Test_SolveSudoku(t *testing.T) {
	type args struct {
		s               []int
		maxValue        int
		maxIndex        int
		relatedElements [][]int
	}
	tests := []struct {
		name string
		args args
		err  error
		want []int
	}{
		{
			name: "sample sudoku #1",
			args: args{
				s: []int{
					0, 0, 3, 0, 2, 0, 6, 0, 0,
					9, 0, 0, 3, 0, 5, 0, 0, 1,
					0, 0, 1, 8, 0, 6, 4, 0, 0,
					0, 0, 8, 1, 0, 2, 9, 0, 0,
					7, 0, 0, 0, 0, 0, 0, 0, 8,
					0, 0, 6, 7, 0, 8, 2, 0, 0,
					0, 0, 2, 6, 0, 9, 5, 0, 0,
					8, 0, 0, 2, 0, 3, 0, 0, 9,
					0, 0, 5, 0, 1, 0, 3, 0, 0,
				},
				maxValue:        9,
				maxIndex:        80,
				relatedElements: autogen.RelatedElements3,
			},
			want: []int{
				4, 8, 3, 9, 2, 1, 6, 5, 7,
				9, 6, 7, 3, 4, 5, 8, 2, 1,
				2, 5, 1, 8, 7, 6, 4, 9, 3,
				5, 4, 8, 1, 3, 2, 9, 7, 6,
				7, 2, 9, 5, 6, 4, 1, 3, 8,
				1, 3, 6, 7, 9, 8, 2, 4, 5,
				3, 7, 2, 6, 8, 9, 5, 1, 4,
				8, 1, 4, 2, 5, 3, 7, 6, 9,
				6, 9, 5, 4, 1, 7, 3, 8, 2,
			},
		},
		{
			name: "sample sudoku #2",
			args: args{
				s: []int{
					2, 0, 0, 0, 8, 0, 3, 0, 0,
					0, 6, 0, 0, 7, 0, 0, 8, 4,
					0, 3, 0, 5, 0, 0, 2, 0, 9,
					0, 0, 0, 1, 0, 5, 4, 0, 8,
					0, 0, 0, 0, 0, 0, 0, 0, 0,
					4, 0, 2, 7, 0, 6, 0, 0, 0,
					3, 0, 1, 0, 0, 7, 0, 4, 0,
					7, 2, 0, 0, 4, 0, 0, 6, 0,
					0, 0, 4, 0, 1, 0, 0, 0, 3,
				},
				maxValue: 9,
				maxIndex: 80,
				//related_elements: ag.related_elements,
			},
			want: []int{
				2, 4, 5, 9, 8, 1, 3, 7, 6,
				1, 6, 9, 2, 7, 3, 5, 8, 4,
				8, 3, 7, 5, 6, 4, 2, 1, 9,
				9, 7, 6, 1, 2, 5, 4, 3, 8,
				5, 1, 3, 4, 9, 8, 6, 2, 7,
				4, 8, 2, 7, 3, 6, 9, 5, 1,
				3, 9, 1, 6, 5, 7, 8, 4, 2,
				7, 2, 8, 3, 4, 9, 1, 6, 5,
				6, 5, 4, 8, 1, 2, 7, 9, 3,
			},
		},
		{
			name: "sample sudoku #3",
			args: args{
				s: []int{
					0, 0, 0, 0, 0, 0, 9, 0, 7,
					0, 0, 0, 4, 2, 0, 1, 8, 0,
					0, 0, 0, 7, 0, 5, 0, 2, 6,
					1, 0, 0, 9, 0, 4, 0, 0, 0,
					0, 5, 0, 0, 0, 0, 0, 4, 0,
					0, 0, 0, 5, 0, 7, 0, 0, 9,
					9, 2, 0, 1, 0, 8, 0, 0, 0,
					0, 3, 4, 0, 5, 9, 0, 0, 0,
					5, 0, 7, 0, 0, 0, 0, 0, 0,
				},
				maxValue: 9,
				maxIndex: 80,
				//related_elements: related_elements,
			},
			want: []int{
				4, 6, 2, 8, 3, 1, 9, 5, 7,
				7, 9, 5, 4, 2, 6, 1, 8, 3,
				3, 8, 1, 7, 9, 5, 4, 2, 6,
				1, 7, 3, 9, 8, 4, 2, 6, 5,
				6, 5, 9, 3, 1, 2, 7, 4, 8,
				2, 4, 8, 5, 6, 7, 3, 1, 9,
				9, 2, 6, 1, 7, 8, 5, 3, 4,
				8, 3, 4, 2, 5, 9, 6, 7, 1,
				5, 1, 7, 6, 4, 3, 8, 9, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveSudoku(tt.args.s, tt.args.maxValue, tt.args.maxIndex, autogen.RelatedElements3)
			if err != nil {
				t.Errorf("err = %v, want %v", err, nil)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}

		})
	}
}
