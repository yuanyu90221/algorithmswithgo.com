package module01

import "fmt"

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {

	for i := 1; i < n; i++ {
		PrintFizzBuzzValue(i)
		fmt.Print(", ")
	}
	PrintFizzBuzzValue(n)
	fmt.Println()
}
func PrintFizzBuzzValue(i int) {
	var is3times uint
	var is5times uint
	if i%3 == 0 {
		is3times = 0x01
	}
	if i%5 == 0 {
		is5times = 0x10
	}
	var judge uint = is3times | is5times
	switch judge {
	case 0x00:
		fmt.Printf("%d", i)
	case 0x01:
		fmt.Printf("Fizz")
	case 0x10:
		fmt.Printf("Buzz")
	case 0x11:
		fmt.Printf("Fizz Buzz")
	}
}
