package avltree

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type IntComparable int

func (i IntComparable) Less(anotherInt IntComparable) bool {
	return i < anotherInt
}

func TestAvlTree(t *testing.T) {
	t.Run("test inserts", func(t *testing.T) {
		elements := make(map[IntComparable]float64)
		for i := 0; i < 1_000_000; i++ {
			elements[IntComparable(rand.Int())] = rand.Float64()
		}

		a := NewAvlTree[IntComparable, float64]()
		for k, v := range elements {
			a.Insert(k, v)
		}

		for k, v := range elements {
			actualV, _ := a.Search(k)
			assert.Equal(t, v, actualV)
		}
	})

	t.Run("test remove", func(t *testing.T) {
		elements := map[IntComparable]float64{
			23:  112.3,
			543: 193.4,
			984: 190.2,
		}

		a := NewAvlTree[IntComparable, float64]()
		for k, v := range elements {
			a.Insert(k, v)
		}

		a.Remove(23)
		delete(elements, 23)

		e, ok := a.Search(23)
		assert.Equal(t, float64(0), e)
		assert.Equal(t, false, ok)

		for k, v := range elements {
			actualV, _ := a.Search(k)
			assert.Equal(t, v, actualV)
		}
	})
}
