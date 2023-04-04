package linkedlist

import (
	"algo-and-ds/iterator"
)

type Iterator[TVal comparable] struct {
	ll *LinkedList[TVal]

	curNode *linkedListNode[TVal]
}

func newIterator[TVal comparable](ll *LinkedList[TVal]) iterator.Iterator[TVal] {
	return &Iterator[TVal]{
		ll: ll,
	}
}

func (i *Iterator[TVal]) Next() bool {
	if i.curNode == nil {
		i.curNode = i.ll.head
		return i.curNode != nil
	}

	if i.curNode.next == nil {
		return false
	}

	i.curNode = i.curNode.next
	return true
}

func (i *Iterator[TVal]) CurrentElement() TVal {
	return i.curNode.val
}
