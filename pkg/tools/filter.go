package tools

import "github.com/jeeo/functional-go/pkg/iterable"

type filterIterator[T any] struct {
	iter      iterable.Iterable[T]
	predicate func(T) bool
}

func (f filterIterator[T]) HasNext() bool {
	return f.iter.HasNext()
}

func (f filterIterator[T]) Next() *T {
	if !f.HasNext() {
		return nil
	}
	next := *f.iter.Next()
	for !f.predicate(next) {
		next = *f.iter.Next()
	}
	return &next
}

func Filter[T any](iter iterable.Iterable[T], f func(T) bool) iterable.Iterable[T] {
	return &filterIterator[T]{
		iter:      iter,
		predicate: f,
	}
}
