package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	// result := ""
	// tokens := strings.Split(word, "")
	// length := len(tokens)
	// if length == 0 {
	// 	return result
	// } else {
	// 	for i := length - 1; i >= 0; i-- {
	// 		result += tokens[i]
	// 	}
	// }
	result := ReverseWithRume(word)
	return result
}

func ReverseWithRume(word string) string {
	var result string
	for _, rume := range word { // Notice character is not necessary contain 1 Bytes
		result = string(rume) + result
	}
	return result
}
