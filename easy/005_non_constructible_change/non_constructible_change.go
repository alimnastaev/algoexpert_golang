package non_constructible_change

import "sort"

/*
! Non-Constructible Change
	Given an array of positive integers representing the values of coins in your possession,
	write a function that returns the minimum amount of change (the minimum sum of money) that you cannot create.
	The given coins can have any positive integer value and aren't necessarily unique
	(i.e., you can have multiple coins of the same value).

	For example, if you're given coins = [1, 2, 5], the minimum amount of change that you can't create is 4.
	If you're given no coins, the minimum amount of change that you can't create is 1.

! Example:
	coins = [5, 7, 1, 1, 2, 3, 22]
	output: 20
*/

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)

	change := 0
	for _, c := range coins {
		if change+1 < c {
			return change + 1
		}
		
		change += c
	}
	return change + 1
}
