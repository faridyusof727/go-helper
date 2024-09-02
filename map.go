package gohelper

func (c *Collection[T]) Map(f func(data T, index int) T) ([]T, error) {
	result := make([]T, len(c.data))
	for i, item := range c.data {
		result[i] = f(item, i)
	}

	return result, nil
}
