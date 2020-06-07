package module01

import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	result := ""
	tokens := strings.Split(word, "")
	length := len(tokens)
	if length == 0 {
		return result
	} else {
		for i := length - 1; i >= 0; i-- {
			result += tokens[i]
		}
	}
	return result
}
