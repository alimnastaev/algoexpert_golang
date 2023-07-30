package run_length_encoding

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RunLength(t *testing.T) {
	input := "AAAAAAAAAAAAABBCCCCDD"
	expected := "9A4A2B4C2D"
	actual := RunLengthEncoding(input)
	require.Equal(t, expected, actual)

	input = "************^^^^^^^$$$$$$%%%%%%%!!!!!!AAAAAAAAAAAAAAAAAAAA"
	expected = "9*3*7^6$7%6!9A9A2A"
	actual = RunLengthEncoding(input)
	require.Equal(t, expected, actual)

	input = "........______=========AAAA   AAABBBB   BBB"
	expected = "8.6_9=4A3 3A4B3 3B"
	actual = RunLengthEncoding(input)
	require.Equal(t, expected, actual)
}
