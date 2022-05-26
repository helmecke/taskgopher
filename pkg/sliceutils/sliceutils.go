package sliceutils

// StrIndexOf get the index of the given value in the given string slice,
// or -1 if not found.
func StrIndexOf(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return -1
}

// StrSliceContains returns true if string is in slice of strings
func StrSliceContains(slice []string, value string) bool {
	return StrIndexOf(slice, value) != -1
}

// IntIndexOf get the index of the given value in the given string slice,
// or -1 if not found.
func IntIndexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return -1
}

// IntSliceContains returns true if string is in slice of strings
func IntSliceContains(slice []int, value int) bool {
	return IntIndexOf(slice, value) != -1
}
