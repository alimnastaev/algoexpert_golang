package caesar_cipher_encryptor

import "fmt"

/*
! Caesar Cipher Encryptor
https://www.algoexpert.io/questions/caesar-cipher-encryptor

Given a non-empty string of lowercase letters and a non-negative integer representing
a key, write a function that returns a new string obtained by shifting every letter
in the input string by k positions in the alphabet, where k is the key.

Note that letters should "wrap" around the alphabet; in other words,
the letter z shifted by one returns the letter a.

! Example:
	string = "xyz"
	key = 2

	output = "zab"
*/

var start = 96
var end = 122

func CaesarCipherEncryptor(str string, key int) (output string) {
	key = key % 26

	for _, c := range str {
		char := int(c) + key
		offset := char - end

		if offset > 0 {
			char = start + offset
		}

		output += fmt.Sprintf("%c", char)
	}

	return
}
