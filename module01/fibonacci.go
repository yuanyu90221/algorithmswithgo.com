package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//
func Fibonacci(n int) int {
	// recursive solution
	// if n < 2 {
	// 	return n
	// } else {
	// 	return Fibonacci(n-1) + Fibonacci(n-2)
	// }
	//iteration
	result := n
	if n < 2 {
		return result
	}
	fn_1 := 0
	fn_2 := 1
	for idx := 2; idx <= n; idx++ {
		result = fn_1 + fn_2
		fn_1 = fn_2
		fn_2 = result
	}
	return result
}
