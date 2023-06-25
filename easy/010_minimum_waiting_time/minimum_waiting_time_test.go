package minimum_waiting_time

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedSquaredArray(t *testing.T) {
	tests := []struct {
		name    string
		queries []int
		want    int
	}{
		{
			name:    "1",
			queries: []int{3, 2, 1, 2, 6},
			want:    17,
		},
		{
			name:    "2",
			queries: []int{17, 4, 3},
			want:    10,
		},
		{
			name:    "3",
			queries: []int{1, 1, 1, 4, 5, 6, 8, 1, 1, 2, 1},
			want:    81,
		},
		{
			name:    "4",
			queries: []int{5, 4, 3, 2, 1},
			want:    20,
		},
		{
			name:    "5",
			queries: []int{8, 9},
			want:    8,
		},
	}

	for _, tt := range tests {
		var got int
		t.Run(tt.name, func(t *testing.T) {
			got = MinimumWaitingTime(tt.queries)
			require.Equal(t, tt.want, got)

			got = MinimumWaitingTime2(tt.queries)
			require.Equal(t, tt.want, got)
		})
	}
}
