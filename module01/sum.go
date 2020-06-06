package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	// sum := 0
	// for _, val := range numbers {
	// 	sum += val
	// }
	// return sum
	return RecusiveSum(numbers)
}

func RecusiveSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	} else {
		return RecusiveSum(numbers[1:]) + numbers[0]
	}
}
