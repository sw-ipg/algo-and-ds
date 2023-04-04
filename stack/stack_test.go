package stack

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Test Stack LIFO Order", func(t *testing.T) {
		elementsReversed := make([]int, 10)
		elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		copy(elementsReversed, elements)
		sort.Slice(elementsReversed, func(i, j int) bool {
			return elementsReversed[j] < elementsReversed[i]
		})

		s := NewStack[int]()
		for _, e := range elements {
			s.Push(e)
		}

		actualElements := make([]int, 0)
		for {
			e, ok := s.Pop()
			if !ok {
				break
			}

			actualElements = append(actualElements, e)
		}

		assert.Equal(t, elementsReversed, actualElements)
	})

	t.Run("test help methods", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(3)
		s.Push(2)

		assert.Equal(t, 2, s.Top())
		assert.Equal(t, 2, s.Size())
		assert.Equal(t, false, s.IsEmpty())

		s.Erase()
		assert.Equal(t, 0, s.Size())
		assert.Equal(t, true, s.IsEmpty())
	})
}
