package first_non_repeating_character

/*
! First Non-Repeating Character
Write a function that takes in a string of lowercase English-alphabet letters and
returns the index of the string's first non-repeating character.

The first non-repeating character is the first character in a string that occurs only once.

If the input string doesn't have any non-repeating characters, your function should return -1.

! Example:
	string = "abcdcaf"
	output: 1 // The first non-repeating character is "b" and is found at index 1.
*/

func FirstNonRepeatingCharacter(str string) int {
	lookupTable := make(map[rune]int)

	for _, char := range str {
		lookupTable[char]++
	}

	for idx, char := range str {
		if lookupTable[char] == 1 {
			return idx
		}
	}

	return -1
}
