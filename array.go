package gohelper

import "reflect"

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
