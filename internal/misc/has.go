package misc

// Has returns true if toFind is in the provided collection (a slice)
func Has[T comparable](collection []T, toFind T) bool {
	for _, thing := range collection {
		if thing == toFind {
			return true
		}
	}
	return false
}

// HasFunc is like [Has], but it accepts a comparison function to test for equivalence
func HasFunc[T any](collection []T, toFind T, isEqual func(a, b T) bool) bool {
	for _, thing := range collection {
		if isEqual(toFind, thing) {
			return true
		}
	}
	return false
}
