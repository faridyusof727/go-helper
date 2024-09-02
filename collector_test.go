package gohelper_test

import (
	"testing"

	gohelper "github.com/faridyusof727/go-helper"
	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Name string
	Age  int
}


func TestCollector(t *testing.T) {
	collector := gohelper.NewCollection[testStruct]()

	data := []testStruct{
		{Name: "John", Age: 20},
		{Name: "Jane", Age: 21},
		{Name: "Doe", Age: 22},
	}

	res, err := collector.Collect(data).Map(func(data testStruct, index int) testStruct {
		return testStruct{
			Name: data.Name + " Yusof",
			Age:  data.Age + 100,
		}
	})
	require.NoError(t, err)
	t.Logf("res: %v", res)
}
