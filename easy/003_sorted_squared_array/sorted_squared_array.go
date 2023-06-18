package sorted_squared_array

/*
! Validate Subsequence
	Given two non-empty arrays of integers, write a function that
	determines whether the second array is a subsequence of the first one.
	A subsequence of an array is a set of numbers that aren't necessarily adjacent in the array,
	but that are in the same order as they appear in the array.
	For instance, the numbers [1, 3, 4] form a subsequence of the array [1, 2, 3, 4] , and so do the numbers [2, 4].
	Note that a single number in an array and the array itself are both valid subsequences of the array.

! Example:
	input: array = [5, 1, 22, 25, 6, -1, 8, 10], sequence = [1, 6, -1, 10]
	output: true
*/

func SortedSquaredArray1(array []int) []int {
	output := make([]int, len(array))

	leftIdx := 0
	left := array[leftIdx]

	rightIdx := len(array) - 1
	right := array[rightIdx]

	idx := len(array) - 1
	for {
		left = abs(left)
		right = abs(right)

		if idx == 0 {
			output[idx] = right * left
			break
		}

		if left > right {
			output[idx] = left * left
			leftIdx++
			left = array[leftIdx]

			idx--
		} else {
			output[idx] = right * right
			rightIdx--
			right = array[rightIdx]

			idx--
		}
	}
	return output
}

func SortedSquaredArray2(array []int) []int {
	output := make([]int, len(array))

	leftIdx := 0
	rightIdx := len(array) - 1

	for idx := len(array) - 1; idx >= 0; idx-- {
		left := abs(array[leftIdx])
		right := abs(array[rightIdx])

		if left > right {
			output[idx] = left * left
			leftIdx++
		} else {
			output[idx] = right * right
			rightIdx--
		}
	}

	return output
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
