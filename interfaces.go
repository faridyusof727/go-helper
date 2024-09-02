package gohelper

type Collector[T any] interface {
	Collect(data []T) Collector[T]

	Get() []T

	Mapper[T]
}

type Mapper[T any] interface {
	Map(f func(data T, index int) T) ([]T, error)
}
