package first_non_repeating_character

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := "abcdcaf"
	expected := 1
	actual := FirstNonRepeatingCharacter(input)
	require.Equal(t, expected, actual)
}
