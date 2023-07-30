package bubble_sort

/*
! Bubble Sort
https://www.algoexpert.io/questions/bubble-sort

Write a function that takes in an array of integers and returns a sorted version of that array.
Use the Bubble Sort algorithm to sort the array.

! Example:
    ascending:
	array = [8, 5, 2, 9, 5, 6, 3]
	output: [2, 3, 5, 5, 6, 8, 9]

    descending:
	array = [8, 5, 2, 9, 5, 6, 3]
	output: [9, 8, 6, 5, 5, 3, 2]
*/

func SortAsc(array []int) []int {
	isSorted := false
	counter := 0

	for !isSorted {
		isSorted = true

		for i := 0; i < len(array)-1-counter; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isSorted = false
			}
		}

		counter++
	}

	return array
}

func SortDesc(array []int) []int {
	isSorted := false
	counter := 0

	for !isSorted {
		isSorted = true

		for i := len(array) - 1; i > counter; i-- {
			if array[i] > array[i-1] {
				array[i-1], array[i] = array[i], array[i-1]
				isSorted = false
			}
		}

		counter++
	}

	return array
}
