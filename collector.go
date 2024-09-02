package gohelper

type Collection[T any] struct {
	data []T
}

func (c *Collection[T]) Collect(data []T) Collector[T] {
	c.data = data
	return c
}

func (c *Collection[T]) Get() []T {
	return c.data
}

func NewCollection[T any]() Collector[T] {
	return &Collection[T]{}
}
