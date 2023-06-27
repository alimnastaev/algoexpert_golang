package find_three_largest_numbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := []int{18, 141, 541}
	output := FindThreeLargestNumbers([]int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7})
	require.Equal(t, expected, output)
}
