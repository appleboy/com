package array

// In check string in array.
func In(needle string, haystack []string) ([]string, bool) {
	if len(haystack) == 0 {
		return haystack, false
	}

	for i, val := range haystack {
		if val == needle {
			haystack = append(haystack[:i], haystack[i+1:]...)
			return haystack, true
		}
	}

	return haystack, false
}

// Diff show difference in two array.
func Diff(s, t []string) []string {
	slice1 := make([]string, len(s))
	slice2 := make([]string, len(t))
	copy(slice1, s)
	copy(slice2, t)
	v := []string{}
	if len(slice1) == 0 && len(slice2) == 0 {
		return []string{}
	}

	if len(slice1) == 0 {
		return slice2
	}

	if len(slice2) == 0 {
		return slice1
	}

	if len(slice1) > len(slice2) {
		slice1, slice2 = slice2, slice1
	}

	for _, val := range slice1 {
		if newT, ok := In(val, slice2); ok {
			slice2 = newT
			continue
		}

		v = append(v, val)
	}

	if len(slice2) > 0 {
		v = append(v, slice2...)
	}

	return v
}
