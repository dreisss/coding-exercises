package main

import "slices"

func smaller_greater(numbers []int) (int, int) {
	smaller := numbers[0]
	greater := numbers[0]

	for n := range numbers {
		if numbers[n] < smaller {
			smaller = numbers[n]
		}

		if numbers[n] > greater {
			greater = numbers[n]
		}
	}

	return smaller, greater
}

func smarter(numbers []int) (int, int) {
	slices.Sort(numbers)

	return numbers[0], numbers[len(numbers)-1]
}
