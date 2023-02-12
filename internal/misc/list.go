package misc

// RemoveFrom returns the same collection with item removed, if item was in it
func RemoveFrom[T comparable](collection []T, item T) []T {
	for idx, itm := range collection {
		if itm == item {
			return append(collection[:idx], collection[idx+1:]...)
		}
	}
	return collection
}
