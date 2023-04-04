package stack

type Stack[TVal any] struct {
	head int
	data []TVal
}

func NewStack[TVal any]() *Stack[TVal] {
	return &Stack[TVal]{
		head: -1,
	}
}

func NewStackWithCapacity[TVal any](cap int) *Stack[TVal] {
	s := NewStack[TVal]()
	s.data = make([]TVal, 0, cap)
	return s
}

func (s *Stack[TVal]) Push(val TVal) {
	s.head++
	s.data = append(s.data, val)
}

func (s *Stack[TVal]) Pop() (val TVal, ok bool) {
	if s.head < 0 {
		return val, false
	}

	val = s.data[s.head]
	s.head--
	return val, true
}

func (s *Stack[TVal]) Top() (val TVal) {
	return s.data[s.head]
}

func (s *Stack[TVal]) Size() int {
	return len(s.data)
}

func (s *Stack[TVal]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[TVal]) Erase() {
	s.head = 0
	s.data = nil
	return
}
