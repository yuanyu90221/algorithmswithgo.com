package module01

import (
	"fmt"
	"math"
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	length := len(value)
	result := 0
	// for idx, value := range value {
	// 	result += ConvertDigit(value, base, length-idx-1)
	// }
	// Sscanf version solution
	multiplier := 1
	for i := length - 1; i >= 0; i-- {
		var val int
		fmt.Sscanf(string(value[i]), "%X", &val)
		result += multiplier * val
		multiplier *= base
	}
	return result
}

func ConvertDigit(digit rune, base int, index int) int {
	switch digit {
	case 'A':
		return int(math.Pow(float64(base), float64(index))) * 10
	case 'B':
		return int(math.Pow(float64(base), float64(index))) * 11
	case 'C':
		return int(math.Pow(float64(base), float64(index))) * 12
	case 'D':
		return int(math.Pow(float64(base), float64(index))) * 13
	case 'E':
		return int(math.Pow(float64(base), float64(index))) * 14
	case 'F':
		return int(math.Pow(float64(base), float64(index))) * 15
	default:
		result, _ := strconv.Atoi(string(digit))
		return int(math.Pow(float64(base), float64(index))) * result
	}
}
