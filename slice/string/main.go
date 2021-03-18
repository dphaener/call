package slice

// Safely fetches an element from a slice of strings. If the index is
// out of bounds, returns an empty string.
func At(slice []string, index int) string {
	if len(slice) < index+1 {
		return ""
	}
	return slice[index]
}

// Safely fetches a range from a slice starting from the given index. Returns
// an empty slice if the index is out of bounds.
func From(slice []string, index int) []string {
	if len(slice) < index+1 {
		var emptySlice []string
		return emptySlice
	}
	return slice[index:]
}

// Safely fetches a range from a slice starting from 0 up to the given index.
// Returns an empty slice if the index is out of bounds.
func To(slice []string, index int) []string {
	if len(slice) < index+1 {
		var emptySlice []string
		return emptySlice
	}
	return slice[:index]
}

// Reverse a slice of strings in place.
func Reverse(slice []string) []string {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}

// Removes any empty strings from a slice of strings.
func Compact(slice []string) (compactedSlice []string) {
	for _, str := range slice {
		if str != "" {
			compactedSlice = append(compactedSlice, str)
		}
	}

	return
}

// Takes a slice of strings and returns a new slice of strings with only
// unique values.
func Uniq(slice []string) (unique []string) {
	for i := range slice {
		if !Contains(slice, slice[i]) {
			unique = append(unique, slice[i])
		}
	}
	return
}

// Iterates through a slice of strings and determines if the given string
// is contained in the slice.
func Contains(slice []string, value string) bool {
	for i := range slice {
		if slice[i] == value {
			return true
		}
	}
	return false
}
