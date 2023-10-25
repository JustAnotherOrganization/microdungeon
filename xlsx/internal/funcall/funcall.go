package funcall

import (
	"reflect"
	"regexp"
)

var functionCallRegex = regexp.MustCompile(`\w(.*)`)

func Function(key string, elem reflect.Value) reflect.Value {
	if !functionCallRegex.MatchString(key) {
		return elem
	}

	elem.Interface()
}
