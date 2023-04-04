package iterator

import "testing"

func CompareWithSlice[T comparable](t *testing.T, i Iterator[T], slice []T) {
	idx := 0
	for i.Next() {
		if len(slice) <= idx {
			t.Fatalf("len(slice) < len(Iterator)")
		}

		if slice[idx] != i.CurrentElement() {
			t.Fatalf("Element at %d (=%v) not equals with iterator value (=%v)", idx, slice[idx], i.CurrentElement())
		}

		idx++
	}

	if len(slice) != idx {
		t.Fatalf("len(slice) > len(Iterator)")
	}
}
