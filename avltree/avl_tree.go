package avltree

type Comparable[T any] interface {
	comparable
	Less(another T) bool
}

type AvlTree[TKey Comparable[TKey], TValue any] struct {
	root *avlNode[TKey, TValue]
}

type avlNode[TKey Comparable[TKey], TValue any] struct {
	key TKey
	val TValue

	height int
	left   *avlNode[TKey, TValue]
	right  *avlNode[TKey, TValue]
}

func NewAvlTree[TKey Comparable[TKey], TValue any]() *AvlTree[TKey, TValue] {
	return &AvlTree[TKey, TValue]{}
}

func (a *AvlTree[TKey, TValue]) Insert(key TKey, val TValue) {
	a.root = a.root.insert(key, val)
}

func (a *AvlTree[TKey, TValue]) Remove(key TKey) {
	a.root = a.root.remove(key)
}

func (a *AvlTree[TKey, TValue]) Search(key TKey) (v TValue, ok bool) {
	n := a.root.search(key)
	if n == nil {
		return v, false
	}

	return n.val, true
}

func (n *avlNode[TKey, TValue]) insert(key TKey, val TValue) *avlNode[TKey, TValue] {
	if n == nil {
		return &avlNode[TKey, TValue]{
			key:    key,
			val:    val,
			height: 1,
		}
	}

	if key.Less(n.key) {
		n.left = n.left.insert(key, val)
	} else if n.key == key {
		n.key = key
	} else {
		n.right = n.right.insert(key, val)
	}

	return n.postChangingBalancing()
}

func (n *avlNode[TKey, TValue]) remove(key TKey) *avlNode[TKey, TValue] {
	if n == nil {
		return nil
	}

	if key.Less(n.key) {
		n.left = n.left.remove(key)
	} else if n.key.Less(key) {
		n.right = n.right.remove(key)
	} else {
		if n.left != nil && n.right != nil {
			minNode := n.right.findMin()
			n.key = minNode.key
			n.right = n.right.remove(key)
		} else if n.left != nil {
			n = n.left
		} else if n.right != nil {
			n = n.right
		} else {
			n = nil
			return n
		}
	}

	return n.postChangingBalancing()
}

func (n *avlNode[TKey, TValue]) search(key TKey) *avlNode[TKey, TValue] {
	if n == nil {
		return nil
	}

	if key.Less(n.key) {
		return n.left.search(key)
	} else if n.key.Less(key) {
		return n.right.search(key)
	} else {
		return n
	}
}

func (n *avlNode[TKey, TValue]) postChangingBalancing() *avlNode[TKey, TValue] {
	if n == nil {
		return nil
	}

	n.calcHeight()
	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}

		return n.rotateLeft()
	} else if balanceFactor == 2 {
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}

		return n.rotateRight()
	}

	return n
}

func (n *avlNode[TKey, TValue]) calcHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *avlNode[TKey, TValue]) getHeight() int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *avlNode[TKey, TValue]) rotateRight() *avlNode[TKey, TValue] {
	nr := n.left
	n.left = nr.right
	nr.right = n

	n.calcHeight()
	nr.calcHeight()
	return nr
}

func (n *avlNode[TKey, TValue]) rotateLeft() *avlNode[TKey, TValue] {
	nr := n.right
	n.right = nr.left
	nr.left = n

	n.calcHeight()
	nr.calcHeight()
	return nr
}

func (n *avlNode[TKey, TValue]) findMin() *avlNode[TKey, TValue] {
	if n.left != nil {
		return n.left.findMin()
	} else {
		return n
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
