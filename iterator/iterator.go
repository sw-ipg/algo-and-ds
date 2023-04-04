package iterator

type Iterator[TVal any] interface {
	Next() bool
	CurrentElement() TVal
}

type Iterable[TVal any] interface {
	Iterator() Iterator[TVal]
}
