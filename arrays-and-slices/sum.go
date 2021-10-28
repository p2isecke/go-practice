package arrays_and_slices

func Sum(numbers []int) int {
	sum := 0

	// range returns two values: index and the value
	// ignore the index value using the blank identifier _
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		// add to a new slice that we'll return
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// Tail: all items in the collection except the first one
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// add to a new slice that we'll return
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums

}