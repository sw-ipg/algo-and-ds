package hashtable

type Hashable[T any] interface {
	comparable
	GetHash() int
}

type HashFunc[T any] func(val T) int

const (
	DefaultCapacity = 100
)

type HashTable[TKey comparable, TValue any] struct {
	nodes []node[TKey, TValue]
	cap   int // общий размер хэш таблицы
	size  int // текущее количество элементов в хэш таблице

	hf HashFunc[TKey]
}

type node[TKey comparable, TValue any] struct {
	key TKey
	val TValue

	next    *node[TKey, TValue] // связанный список для разрешения коллизий
	notZero bool
}

const (
	_maxLoadFactor = 0.8 // максимальный фактор загрузки таблицы, после чего происходит growth
	_growthFactor  = 2   // фактор роста хэш таблицы
)

func NewHashTable[TKey Hashable[TKey], TValue any](cap int) *HashTable[TKey, TValue] {
	return &HashTable[TKey, TValue]{
		nodes: make([]node[TKey, TValue], cap),
		cap:   cap,
		hf: func(val TKey) int {
			return val.GetHash()
		},
	}
}

func NewHashTableWithHashFunc[TKey comparable, TValue any](cap int, hashFunc HashFunc[TKey]) *HashTable[TKey, TValue] {
	return &HashTable[TKey, TValue]{
		nodes: make([]node[TKey, TValue], cap),
		cap:   cap,
		hf:    hashFunc,
	}
}

func (h *HashTable[TKey, TValue]) Insert(key TKey, val TValue) {
	if h.loadFactor() > _maxLoadFactor {
		h.growth()
	}

	n := node[TKey, TValue]{key: key, val: val, notZero: true}
	idx := h.getIndex(key)
	if h.nodes[idx].notZero {
		if h.nodes[idx].key != key { // коллизия
			var prev *node[TKey, TValue]
			cur := &h.nodes[idx]
			for cur != nil {
				prev = cur
				cur = cur.next
			}

			prev.next = &n
		} else {
			h.nodes[idx].val = val
			return
		}
	} else {
		h.nodes[idx] = n
	}

	h.size++
}

func (h *HashTable[TKey, TValue]) Remove(key TKey) {
	idx := h.getIndex(key)
	var prev *node[TKey, TValue]
	cur := &h.nodes[idx]
	for cur != nil {
		if cur.key == key && cur.notZero {
			if prev == nil {
				if cur.next == nil {
					h.nodes[idx].notZero = false
				} else {
					h.nodes[idx] = *cur.next
				}

				break
			}

			prev.next = cur.next
			break
		}

		prev = cur
		cur = cur.next
	}
}

func (h *HashTable[TKey, TValue]) Get(key TKey) (v TValue, ok bool) {
	idx := h.getIndex(key)
	if h.nodes[idx].key == key && h.nodes[idx].notZero { // fast-path
		return h.nodes[idx].val, true
	}

	cur := &h.nodes[idx]
	for cur != nil {
		if cur.key == key && cur.notZero {
			return cur.val, true
		}

		cur = cur.next
	}

	return v, false
}

func (h *HashTable[TKey, TValue]) Size() int {
	return h.size
}

func (h *HashTable[TKey, TValue]) growth() {
	oldNodes := h.nodes
	h.cap = 2 * h.cap
	h.size = 0
	h.nodes = make([]node[TKey, TValue], h.cap)

	for _, oldNode := range oldNodes {
		if oldNode.notZero {
			cur := &oldNode
			for cur != nil {
				h.Insert(cur.key, cur.val)
				cur = cur.next
			}
		}
	}
}

func (h *HashTable[TKey, TValue]) loadFactor() float64 {
	return float64(h.size) / float64(h.cap)
}

func (h *HashTable[TKey, TValue]) getIndex(key TKey) int {
	return h.hf(key) % h.cap
}
