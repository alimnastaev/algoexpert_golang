package sorted_squared_array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedSquaredArray(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "1",
			input: []int{2, 3, 5, 6, 8, 9},
			want:  []int{4, 9, 25, 36, 64, 81},
		},
		{
			name:  "2",
			input: []int{-4, -1},
			want:  []int{1, 16},
		},
		{
			name:  "3",
			input: []int{-7, -3, 1, 9, 22, 30},
			want:  []int{1, 9, 49, 81, 484, 900},
		},
	}

	for _, tt := range tests {
		got := []int{}
		t.Run(tt.name, func(t *testing.T) {
			got = SortedSquaredArray1(tt.input)
			require.Equal(t, got, tt.want)

			got = SortedSquaredArray2(tt.input)
			require.Equal(t, got, tt.want)
		})
	}
}
