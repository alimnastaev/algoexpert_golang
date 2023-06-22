package non_constructible_change

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedSquaredArray(t *testing.T) {
	tests := []struct {
		name  string
		coins []int
		want  int
	}{
		{
			name:  "1",
			coins: []int{5, 7, 1, 1, 2, 3, 22},
			want:  20,
		},
		{
			name:  "2",
			coins: []int{1, 1},
			want:  3,
		},
		{
			name:  "3",
			coins: []int{1, 2, 3, 4, 5, 6, 7},
			want:  29,
		},
		{
			name:  "4",
			coins: []int{1, 1, 1, 1, 1},
			want:  6,
		},
		{
			name:  "5",
			coins: []int{1, 5, 1, 1, 1, 10, 15, 20, 100},
			want:  55,
		},
	}

	var got int

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got = NonConstructibleChange(tt.coins)
			require.Equal(t, tt.want, got)
		})
	}
}
