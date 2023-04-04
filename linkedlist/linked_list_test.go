package linkedlist

import (
	"algo-and-ds/iterator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("test insert", func(t *testing.T) {
		elements := []int{42, 553, 65, 12, 98, 109}
		l := NewLinkedList[int]()

		for _, e := range elements {
			l.Insert(e)
		}

		iterator.CompareWithSlice(t, l.Iterator(), elements)
		assert.Equal(t, len(elements), l.Size())
	})

	t.Run("test remove", func(t *testing.T) {
		elements := []int{42, 552, 90, 440, 214, 999}
		elementsWithout42 := []int{552, 90, 440, 214, 999}

		l := NewLinkedList[int]()
		for _, e := range elements {
			l.Insert(e)
		}

		l.Remove(42)

		iterator.CompareWithSlice(t, l.Iterator(), elementsWithout42)
		assert.Equal(t, len(elementsWithout42), l.Size())
	})

	t.Run("test erase", func(t *testing.T) {
		l := NewLinkedList[int]()
		l.Insert(42)
		l.Insert(54)
		assert.Equal(t, 2, l.Size())
		iterator.CompareWithSlice(t, l.Iterator(), []int{42, 54})

		l.Erase()
		assert.Equal(t, 0, l.Size())
		iterator.CompareWithSlice(t, l.Iterator(), []int{})
	})
}
