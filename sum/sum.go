package sum

// Ints sums up numbers with array of number as input. (public impl)
func Ints(numbers ...int) int {
	return ints(numbers)
}

// Ints sums up numbers with slice of number as input. (private impl)
func ints(numbers []int) int {
	// Return 0 when no input available
	if len(numbers) == 0 {
		return 0
	}

	// Return sums of numbers from index 1 etc with index 0.
	return ints(numbers[1:]) + numbers[0]
}
