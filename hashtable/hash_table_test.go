package hashtable

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestHashTable(t *testing.T) {
	t.Run("test insert", func(t *testing.T) {
		elements := map[Int64Hashable]string{
			42:  "A",
			64:  "B",
			99:  "C",
			109: "D",
		}

		ht := NewHashTable[Int64Hashable, string](DefaultCapacity)
		for k, v := range elements {
			ht.Insert(k, v)
		}

		for k, v := range elements {
			htEl, ok := ht.Get(k)
			if !ok {
				t.Fatalf("not value by key %v", k)
			}

			assert.Equal(t, v, htEl)
		}
	})

	t.Run("test removing", func(t *testing.T) {
		elements := map[Int64Hashable]string{
			42:  "A",
			64:  "B",
			99:  "C",
			109: "D",
		}

		ht := NewHashTable[Int64Hashable, string](DefaultCapacity)
		for k, v := range elements {
			ht.Insert(k, v)
		}

		ht.Remove(42)
		htEl, ok := ht.Get(42)
		assert.Equal(t, false, ok)
		assert.Equal(t, "", htEl)
	})

	t.Run("test growth", func(t *testing.T) {
		elements := map[Int64Hashable]string{
			42:  "A",
			64:  "B",
			99:  "C",
			109: "D",
			553: "E",
			857: "F",
			10:  "G",
		}

		ht := NewHashTable[Int64Hashable, string](4)
		for k, v := range elements {
			ht.Insert(k, v)
		}

		for k, v := range elements {
			htEl, ok := ht.Get(k)
			if !ok {
				t.Fatalf("cannot get htEl for key %v", k)
			}

			assert.Equal(t, v, htEl)
		}

		assert.Equal(t, len(elements), ht.Size())

	})

	t.Run("test collisions", func(t *testing.T) {
		originalTableTimeStart := time.Now()
		elements := make(map[Int64Hashable]float64)
		for i := 0; i < 1_000_000; i++ {
			elements[Int64Hashable(rand.Int())] = rand.Float64()
		}
		originalTableTimeEnd := time.Now()

		myAwesomeHashTableTimeStart := time.Now()
		ht := NewHashTable[Int64Hashable, float64](DefaultCapacity)
		for k, v := range elements {
			ht.Insert(k, v)
		}
		myAwesomeHashTableTimeEnd := time.Now()

		for k, v := range elements {
			htEl, ok := ht.Get(k)
			if !ok {
				t.Fatalf("cannot get htEl for key %v", k)
			}

			assert.Equal(t, v, htEl)
		}

		assert.Equal(t, len(elements), ht.Size())

		log.Printf("original go hashtable insertion time: %s, my awesome hashtable: %s", originalTableTimeEnd.Sub(originalTableTimeStart), myAwesomeHashTableTimeEnd.Sub(myAwesomeHashTableTimeStart))
	})
}
