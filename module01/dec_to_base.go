package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	result := ""
	remain := dec % base
	carry := dec / base
	lastCarry := 0
	for ; carry != 0 || remain != 0; remain = lastCarry % base {
		result = Format(remain) + result
		lastCarry = carry
		carry = carry / base
	}
	fmt.Println(result)
	return result
}

func Format(digit int) string {
	switch digit {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	default:
		return fmt.Sprintf("%d", digit)
	}
}
