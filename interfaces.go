package gohelper

import "golang.org/x/exp/constraints"

type Collector[T interface {
	constraints.Ordered
	comparable
}] interface {
	Collect(data []T) Collector[T]

	Get() []T

	Positioner[T]
	Mapper[T]
}

type Mapper[T comparable] interface {
	Map(f func(data T, index int) T) ([]T, error)
}

type Positioner[T constraints.Ordered] interface {
	After(d T) (T, error)
	Before(d T) (T, error)
}
