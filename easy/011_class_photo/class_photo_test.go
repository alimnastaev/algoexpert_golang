package class_photo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedSquaredArray(t *testing.T) {
	tests := []struct {
		name             string
		redShirtHeights  []int
		blueShirtHeights []int
		want             bool
	}{
		{
			name:             "1",
			redShirtHeights:  []int{5, 8, 1, 3, 4},
			blueShirtHeights: []int{6, 9, 2, 4, 5},
			want:             true,
		},
		{
			name:             "2",
			redShirtHeights:  []int{5, 4},
			blueShirtHeights: []int{5, 6},
			want:             true,
		},
		{
			name:             "2",
			redShirtHeights:  []int{5, 8, 1, 3, 4},
			blueShirtHeights: []int{6, 9, 2, 4, 5},
			want:             true,
		},
		{
			name:             "2",
			redShirtHeights:  []int{5, 6, 7, 2, 3, 1, 2, 3},
			blueShirtHeights: []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:             false,
		},
		{
			name:             "2",
			redShirtHeights:  []int{126},
			blueShirtHeights: []int{125},
			want:             true,
		},
	}

	for _, tt := range tests {
		var got bool
		t.Run(tt.name, func(t *testing.T) {
			got = ClassPhotos(tt.redShirtHeights, tt.blueShirtHeights)
			require.Equal(t, tt.want, got)
		})
	}
}
