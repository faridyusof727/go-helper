package gohelper

import "golang.org/x/exp/constraints"

type Collection[T interface {
	constraints.Ordered
	comparable
}] struct {
	data []T
}

func (c *Collection[T]) Collect(data []T) Collector[T] {
	c.data = data
	return c
}

func (c *Collection[T]) Get() []T {
	return c.data
}

func NewCollection[T interface {
	constraints.Ordered
	comparable
}]() Collector[T] {
	return &Collection[T]{}
}
