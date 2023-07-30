package generate_document

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RunLength(t *testing.T) {
	characters := "Bste!hetsi ogEAxpelrt x "
	document := "AlgoExpert is the Best!"
	expected := true
	actual := GenerateDocument(characters, document)
	require.Equal(t, expected, actual)

	characters = "a hsgalhsa sanbjksbdkjba kjx"
	document = ""
	expected = true
	actual = GenerateDocument(characters, document)
	require.Equal(t, expected, actual)

	characters = "aheaolabbhb"
	document = "hello"
	expected = false
	actual = GenerateDocument(characters, document)
	require.Equal(t, expected, actual)
}
