package caesar_cipher_encryptor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RunLength(t *testing.T) {
	str := "xyz"
	key := 2
	expected := "zab"
	actual := CaesarCipherEncryptor(str, key)
	require.Equal(t, expected, actual)

	str = "abc"
	key = 57
	expected = "fgh"
	actual = CaesarCipherEncryptor(str, key)
	require.Equal(t, expected, actual)
}
