package hashtable

type Int64Hashable int64

func (k Int64Hashable) GetHash() int {
	return int(k ^ (k >> 32)) // Java long hashcode
}
