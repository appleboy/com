package array

import (
	"reflect"
)

// InMap check index in Map
func InMap(needle string, haystack []string) bool {
	newStack := map[string]struct{}{}
	for _, val := range haystack {
		newStack[val] = struct{}{}
	}

	if _, ok := newStack[needle]; ok {
		return true
	}

	return false
}

// InSlice check index in slice
func InSlice(needle string, haystack []string) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}

	return false
}

// InSliceInt64 check index in slice int64
func InSliceInt64(needle int64, haystack []int64) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}

	return false
}

// InArray index in array for interface{}
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// In check string in array.
func In(needle string, haystack []string) ([]string, bool) {
	newHaystack := make([]string, len(haystack))
	copy(newHaystack, haystack)

	if len(newHaystack) == 0 {
		return newHaystack, false
	}

	for i, val := range newHaystack {
		if val == needle {
			newHaystack = append(newHaystack[:i], newHaystack[i+1:]...)
			return newHaystack, true
		}
	}

	return newHaystack, false
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
