package gohelper

import (
	"regexp"
	"strconv"
)

func NumericOnly(input string) (float64, error) {
	var re = regexp.MustCompile("([^.0-9])()")
	return strconv.ParseFloat(re.ReplaceAllString(input, "$2"), 64)
}
