package find_three_largest_numbers

import "math"

/*
! Find Three Largest Numbers
Write a function that takes in an array of at least three integers and, without sorting the input array,
returns a sorted array of the three largest integers in the input array.

The function should return duplicate integers if necessary;
for example, it should return [10, 10, 12] for an input array of [10, 5, 9, 10, 12].

! Example:
	array = [141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7]
	output: [18, 141, 541]
*/

func FindThreeLargestNumbers(array []int) []int {
	output := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	for _, n := range array {
		if n > output[2] {
			shiftAndUpdate(output, n, 2)
		} else if n > output[1] {
			shiftAndUpdate(output, n, 1)
		} else if n > output[0] {
			shiftAndUpdate(output, n, 0)
		}

	}
	return output
}

func shiftAndUpdate(output []int, n, idx int) {
	for i := 0; i < idx+1; i++ {
		if i == idx {
			output[i] = n
		} else {
			output[i] = output[i+1]
		}
	}
}
