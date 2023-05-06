package validate_subsequence

func IsValidSubsequence(array []int, sequence []int) bool {
	seqSize := len(sequence)
	counter := 0
	for _, i := range array {
		if counter < seqSize && i == sequence[counter] {
			counter++
		}
	}

	return counter == seqSize
}
