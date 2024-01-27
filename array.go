package gohelper

import "reflect"
// test
func ArrayIsAccessible(v interface{}) bool {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	case reflect.Array:
		return true
	default:
		return false
	}
}

func ArrayAdd[T any](data []T, s T) []T {
	newData := append(data, s)
	return newData
}

func ArrayAddPt[T any](data *[]T, s T) {
	newData := append(*data, s)
	*data = newData
}
