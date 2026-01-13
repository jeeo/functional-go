package tools

import "github.com/jeeo/functional-go/pkg/iterable"

type mapIterator[T, U any] struct {
	iter   iterable.Iterable[T]
	mapper func(T) *U
}

func (m mapIterator[T, U]) HasNext() bool {
	return m.iter.HasNext()
}

func (m mapIterator[T, U]) Next() *U {
	if !m.HasNext() {
		return nil
	}

	next := m.iter.Next()
	return m.mapper(*next)
}

func Map[T any, U any](iterator iterable.Iterable[T], f func(T) *U) iterable.Iterable[U] {
	return mapIterator[T, U]{
		iter:   iterator,
		mapper: f,
	}
}
