package misc

import (
	"testing"
)

func TestRemoveFrom(t *testing.T) {
	t.Run("nil list", func(t *testing.T) {
		var collection []int
		var item int = 0
		result := RemoveFrom(collection, item)
		if Has(result, item) {
			t.Error("should be absent")
		}
	})
	t.Run("simple list", func(t *testing.T) {
		var collection []int = []int{0, 1, 2, 3}
		var item int = 0
		result := RemoveFrom(collection, item)
		if Has(result, item) {
			t.Error("should be absent", result)
		}
	})
	t.Run("short list", func(t *testing.T) {
		var collection []int = []int{0}
		var item int = 0
		result := RemoveFrom(collection, item)
		if Has(result, item) {
			t.Error("should be absent", result)
		}
	})
	t.Run("last in list", func(t *testing.T) {
		var collection []int = []int{3, 2, 1, 0}
		var item int = 0
		result := RemoveFrom(collection, item)
		if Has(result, item) {
			t.Error("should be absent", result)
		}
	})
	t.Run("middle in list", func(t *testing.T) {
		var collection []string = []string{"a", "b", "c", "d", "e", "f", "g"}
		var item string = "d"
		result := RemoveFrom(collection, item)
		if Has(result, item) {
			t.Error("should be absent", result)
		}
		for _, other := range []string{"a", "b", "c", "e", "f", "g"} {
			if !Has(result, other) {
				t.Error(other, "should be in", result)
			}
		}
	})
}
