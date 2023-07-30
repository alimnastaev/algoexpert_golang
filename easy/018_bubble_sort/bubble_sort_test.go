package bubble_sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := []int{8, 5, 2, 9, 5, 6, 3}
	expected := []int{2, 3, 5, 5, 6, 8, 9}
	actual := SortAsc(input)
	require.Equal(t, expected, actual)

	expected = []int{9, 8, 6, 5, 5, 3, 2}
	actual = SortDesc(input)
	require.Equal(t, expected, actual)
}