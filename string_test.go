package gohelper_test

import (
	"testing"

	gohelper "github.com/faridyusof727/go-helper/v2"
)

func TestNumericOnlyCase1(t *testing.T) {
	numC1, err := gohelper.NumericOnly("MYR 2,400.20")
	if err != nil {
		t.Error("something went wrong")
	}

	if numC1 != 2400.2 {
		t.Error("assertion error")
	}
}

func TestNumericOnlyCase2(t *testing.T) {
	numC1, err := gohelper.NumericOnly("MYR 1,112,400.20")
	if err != nil {
		t.Error("something went wrong")
	}

	if numC1 != 1112400.20 {
		t.Error("assertion error")
	}
}
