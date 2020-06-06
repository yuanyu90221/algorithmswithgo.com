package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	// loop in list to check
	// Notice first argument get from range is idx and seconde is the value
	for _, i := range list {
		if i == num {
			return true
		}
	}
	return false
}
