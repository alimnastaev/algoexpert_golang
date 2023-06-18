package sorted_squared_array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	array := []int{2, 3, 5, 6, 8, 9}
	sequence := []int{4, 9, 25, 36, 64, 81}
	require.Equal(t, SortedSquaredArray2(array), sequence)
}

func TestCase2(t *testing.T) {
	array := []int{-4, -1}
	sequence := []int{1, 16}
	require.Equal(t, SortedSquaredArray2(array), sequence)
}

func TestCase3(t *testing.T) {
	expected := []int{1, 9, 49, 81, 484, 900}
	actual := SortedSquaredArray2([]int{-7, -3, 1, 9, 22, 30})
	require.Equal(t, expected, actual)
}
