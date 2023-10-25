package xlsx

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type TextCell struct {
	cell

	mux    sync.Mutex
	values []Value
	format string
}

func NewTextCell(location, key string, defaultVal interface{}) *TextCell {
	return &TextCell{
		cell: cell{
			location: location,
		},
		values: []Value{
			{
				Key:   key,
				Value: defaultVal,
			},
		},
	}
}

func NewCompositeTextCell(location string, values ...Value) *TextCell {
	return &TextCell{
		cell: cell{
			location: location,
		},
		values: values,
	}
}

func NewCompositeTextCellf(location, format string, values ...Value) *TextCell {
	return &TextCell{
		cell: cell{
			location: location,
		},
		values: values,
		format: format,
	}
}

func (c *TextCell) Location() string {
	return c.location
}

func (c *TextCell) Encode(elem reflect.Value) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for idx, value := range c.values {
		var field = elem

		if strings.Contains(value.Key, ".") {
			for _, fieldName := range strings.Split(value.Key, ".") {
				field = field.FieldByName(fieldName)
			}
		} else {
			field = field.FieldByName(value.Key)
		}

		if field.IsValid() && !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			var to reflect.Type
			if value.Value != nil {
				to = reflect.TypeOf(value.Value)
			} else {
				to = field.Type()
			}

			nv := reflect.New(to)
			nv.Elem().Set(field)
			c.values[idx].Value = nv.Elem().Interface()
		}
	}
}

func (c *TextCell) Decode(elem reflect.Value) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for _, value := range c.values {
		if value.Value == nil {
			// No default provided and not set in original object, skip.
			continue
		}

		field := elem.FieldByName(value.Key)
		if field.IsValid() && field.CanSet() {
			fieldType := field.Type()
			valType := reflect.TypeOf(value.Value)

			if valType.ConvertibleTo(fieldType) {
				field.Set(reflect.ValueOf(value.Value).Convert(fieldType))
			}
		}
	}
}

func (c *TextCell) Interface() interface{} {
	return c.String()
}

func (c *TextCell) String() string {
	//  FIXME: panic if called before Encode.

	c.mux.Lock()
	defer c.mux.Unlock()

	var vals []interface{}
	for _, value := range c.values {
		vals = append(vals, value.Value)
	}

	if c.format != "" {
		return fmt.Sprintf(c.format, vals...)
	}

	strs := make([]string, len(vals))
	for idx, val := range vals {
		switch v := val.(type) {
		case string:
			strs[idx] = v
		case int:
			strs[idx] = strconv.FormatInt(int64(v), base10)
		case int32:
			strs[idx] = strconv.FormatInt(int64(v), base10)
		case int64:
			strs[idx] = strconv.FormatInt(v, base10)
		case uint:
			strs[idx] = strconv.FormatUint(uint64(v), base10)
		case uint32:
			strs[idx] = strconv.FormatUint(uint64(v), base10)
		case uint64:
			strs[idx] = strconv.FormatUint(v, base10)
		case float32:
			strs[idx] = strconv.FormatFloat(float64(v), base10, 6, bit32)
		case float64:
			strs[idx] = strconv.FormatFloat(v, base10, 6, bit32)
		case bool:
			strs[idx] = strconv.FormatBool(v)
		default:
			strs[idx] = fmt.Sprintf("%v", val)
		}
	}

	return strings.Join(strs, ",")
}
