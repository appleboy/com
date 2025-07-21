package array

// Contains checks if a given key of any comparable type exists within a slice.
// It returns true if the key is found, otherwise it returns false.
//
// Time complexity: O(n), where n is the length of the slice.
// Suitable for small to medium slices or infrequent lookups.
// For large slices with frequent lookups, consider using a map for better performance.
//
// Parameters:
//   - slice: A slice of any comparable type T.
//   - key: An element of type T to search for within the slice.
//
// Returns:
//   - bool: True if the key is found in the slice, false otherwise.
func Contains[T comparable](slice []T, key T) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}
