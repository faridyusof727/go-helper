package gohelper

import (
	"errors"
	"slices"
)

func (c *Collection[T]) After(d T) (T, error) {
	var zero T

	slices.Sort(c.data)

	for _, item := range c.data {
		if d < item {
			return item, nil
		}
	}

	return zero, errors.New("not found")
}

func (c *Collection[T]) Before(d T) (T, error) {
	var zero T

	slices.Sort(c.data)

	for i, j := 0, len(c.data)-1; i < j; i, j = i+1, j-1 {
		c.data[i], c.data[j] = c.data[j], c.data[i]
	}

	for _, item := range c.data {
		if d > item {
			return item, nil
		}
	}

	return zero, errors.New("not found")
}
