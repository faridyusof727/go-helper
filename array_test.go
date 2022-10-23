package gohelper_test

import (
	"testing"

	gohelper "github.com/faridyusof727/go-helper/v2"
)

func TestIsAccessibleArray(t *testing.T) {
	test := map[string]interface{}{
		"array": [6]int{2, 3, 5, 7, 11, 13},
	}

	ok := gohelper.ArrayIsAccessible(test["array"])

	if ok == false {
		t.Error("test['array'] is not returning correct assert")
	}
}

func TestIsAccessibleArrayNeg(t *testing.T) {
	test := map[string]interface{}{
		"array": 2,
	}

	ok := gohelper.ArrayIsAccessible(test["array"])

	if ok != false {
		t.Error("test['array'] is not returning correct assert")
	}
}

func TestIsAccessibleSlice(t *testing.T) {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	test := map[string]interface{}{
		"array": s,
	}

	ok := gohelper.ArrayIsAccessible(test["array"])

	if ok == false {
		t.Error("test['array'] is not returning correct assert")
	}
}

func TestAddInteger(t *testing.T) {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	newData := gohelper.ArrayAdd(primes[:], 20)

	if newData[6] != 20 {
		t.Error("assert fail")
	}
}

func TestAddString(t *testing.T) {
	text := []string{"I", "love", "you"}

	newData := gohelper.ArrayAdd(text, "Gordon")

	if newData[3] != "Gordon" {
		t.Error("assert fail")
	}
}

func TestAddStringPt(t *testing.T) {
	text := []string{"I", "love", "you"}

	gohelper.ArrayAddPt(&text, "Gordon")

	if text[3] != "Gordon" {
		t.Error("assert fail")
	}
}
