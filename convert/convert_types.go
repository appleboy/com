package convert

// SliceToPtrSlice converts a slice of values ([]T) to a slice of pointers ([]*T).
// Each element in the result points to the corresponding element in the input slice.
// Note: The pointers are valid only as long as the input slice is not modified.
func SliceToPtrSlice[T any](src []T) []*T {
	dst := make([]*T, len(src))
	for i := range src {
		dst[i] = &src[i]
	}
	return dst
}

// PtrSliceToSlice converts a slice of pointers ([]*T) to a slice of values ([]T).
// If an element in the input slice is nil, the zero value of T is used.
func PtrSliceToSlice[T any](src []*T) []T {
	dst := make([]T, len(src))
	for i, v := range src {
		if v != nil {
			dst[i] = *v
		}
	}
	return dst
}

// MapToPtrMap converts a map of values (map[string]T) to a map of pointers (map[string]*T).
// Each value in the result points to a copy of the corresponding value in the input map.
func MapToPtrMap[T any](src map[string]T) map[string]*T {
	dst := make(map[string]*T, len(src))
	for k, v := range src {
		val := v
		dst[k] = &val
	}
	return dst
}

// PtrMapToMap converts a map of pointers (map[string]*T) to a map of values (map[string]T).
// If a value in the input map is nil, the key is omitted in the result.
func PtrMapToMap[T any](src map[string]*T) map[string]T {
	dst := make(map[string]T, len(src))
	for k, v := range src {
		if v != nil {
			dst[k] = *v
		}
	}
	return dst
}
