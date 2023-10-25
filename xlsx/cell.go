package xlsx

import (
	"reflect"
)

type (
	Cell interface {
		Location() string
		Encode(elem reflect.Value)
		Decode(elem reflect.Value)
		Interface() interface{}
	}

	cell struct {
		location string
	}

	Value struct {
		Key   string
		Value interface{}
	}
)
