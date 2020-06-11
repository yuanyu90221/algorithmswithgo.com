package module01

func GCD(a, b int) int {
	// if b != 0 {
	// 	return GCD(b, a%b)
	// } else {
	// 	return a
	// }
	dividee := a
	divider := b
	for dividee%divider != 0 {
		dividee, divider = divider, dividee%divider
	}
	return divider
}
