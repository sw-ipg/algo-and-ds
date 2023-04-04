package linkedlist

import "algo-and-ds/iterator"

type LinkedList[TVal comparable] struct {
	head *linkedListNode[TVal]
	tail *linkedListNode[TVal]
	size int
}

type linkedListNode[TVal comparable] struct {
	val  TVal
	next *linkedListNode[TVal]
}

func NewLinkedList[TVal comparable]() *LinkedList[TVal] {
	return &LinkedList[TVal]{}
}

func (l *LinkedList[TVal]) Insert(val TVal) {
	newNode := &linkedListNode[TVal]{
		val: val,
	}

	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}

	l.tail = newNode
	l.size++
	return
}

func (l *LinkedList[TVal]) Remove(val TVal) {
	var current, previous *linkedListNode[TVal]

	previous = nil
	current = l.head
	for current != nil {
		if current.val == val {
			if previous == nil {
				l.head = current.next
			} else {
				previous.next = current.next
				if current.next == nil {
					l.tail = previous
				}
			}

			l.size--
			return
		}

		previous = current
		current = current.next
	}
}

func (l *LinkedList[TVal]) Erase() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList[TVal]) Size() int {
	return l.size
}

func (l *LinkedList[TVal]) Iterator() iterator.Iterator[TVal] {
	return newIterator[TVal](l)
}
