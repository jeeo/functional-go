package iterable

type Iterable[T any] interface {
	HasNext() bool
	Next() *T
}

type Iterator[T any] struct {
	array []T
	index int
}

func NewIterable[T any](list []T) *Iterator[T] {
	return &Iterator[T]{
		array: list,
	}
}

func (i *Iterator[T]) HasNext() bool {
	return i.index < len(i.array)
}

func (i *Iterator[T]) Next() *T {
	var item T
	if !i.HasNext() {
		return nil
	}
	item = i.array[i.index]
	i.index++
	return &item
}
