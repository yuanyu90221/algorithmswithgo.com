package module01

import "fmt"

// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//   BaseToBase("E", 16, 2) => "1110"
//
func BaseToBase(value string, base, newBase int) string {
	return _DecToBase(_BaseToDec(value, base), newBase)
}

func _BaseToDec(value string, base int) int {
	result := 0
	multiplier := 1
	for i := len(value) - 1; i >= 0; i-- {
		var digit int
		fmt.Sscanf(string(value[i]), "%X", &digit)
		result += digit * multiplier
		multiplier *= base
	}
	return result
}

func _DecToBase(dec int, base int) string {
	result := ""
	for dec > 0 {
		remain := dec % base
		result = fmt.Sprintf("%X%s", remain, result)
		dec = dec / base
	}
	return result
}
