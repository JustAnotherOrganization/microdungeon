package xlsx_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"justanother.org/microdungeon/xlsx"
)

func TestTextCell_Encode(t *testing.T) {
	for name, tc := range map[string]struct {
		cell     *xlsx.TextCell
		testOb   testOb
		expected string
	}{
		"string": {
			cell: xlsx.NewTextCell("A1", "String", ""),
			testOb: testOb{
				String: "foobar",
			},
			expected: "foobar",
		},
		"composite string (with format)": {
			cell: xlsx.NewCompositeTextCellf(
				"A1",
				"%s %d",
				xlsx.Value{
					Key: "String",
				},
				xlsx.Value{
					Key: "Int",
				},
			),
			testOb: testOb{
				String: "foobar",
				Int:    42,
			},
			expected: "foobar 42",
		},
		"composite string (without format)": {
			cell: xlsx.NewCompositeTextCell(
				"A1",
				xlsx.Value{
					Key: "String",
				},
				xlsx.Value{
					Key: "Int",
				},
			),
			testOb: testOb{
				String: "foobar",
				Int:    42,
			},
			expected: "foobar,42",
		},
		"encode favors the object over provided defaults": {
			cell: xlsx.NewTextCell("A1", "String", "foobaz"),
			testOb: testOb{
				String: "foobar",
			},
			expected: "foobar",
		},
		"encode uses the default if no value is set on the object": {
			cell:     xlsx.NewTextCell("A1", "String", "foobar"),
			testOb:   testOb{},
			expected: "foobar",
		},
		"no key set, cell is for sheet only and does not correlate with the object": {
			cell:     xlsx.NewTextCell("A1", "", "foobar"),
			testOb:   testOb{},
			expected: "foobar",
		},
		"nested string": {
			cell: xlsx.NewTextCell("A1", "Nested.String", ""),
			testOb: testOb{
				Nested: nested{
					String: "foobar",
				},
			},
			expected: "foobar",
		},
	} {
		t.Run(name, func(tt *testing.T) {
			elem := reflect.ValueOf(tc.testOb)
			tc.cell.Encode(elem)
			assert.Equal(tt, tc.expected, tc.cell.String())
		})
	}
}

func TestTextCell_Decode(t *testing.T) {
	for name, tc := range map[string]struct {
		cell     *xlsx.TextCell
		expected testOb
	}{
		"string": {
			cell: xlsx.NewTextCell("A1", "String", "foobar"),
			expected: testOb{
				String: "foobar",
			},
		},
		"composite string (with format)": {
			cell: xlsx.NewCompositeTextCellf(
				"A1",
				"%s %d",
				xlsx.Value{
					Key:   "String",
					Value: "foobar",
				},
				xlsx.Value{
					Key:   "Int",
					Value: 42,
				},
			),
			expected: testOb{
				String: "foobar",
				Int:    42,
			},
		},
		"composite string (without format)": {
			cell: xlsx.NewCompositeTextCell(
				"A1",
				xlsx.Value{
					Key:   "String",
					Value: "foobar",
				},
				xlsx.Value{
					Key:   "Int",
					Value: 42,
				},
			),
			expected: testOb{
				String: "foobar",
				Int:    42,
			},
		},
		"don't panic if value is nil after encode": {
			cell: xlsx.NewCompositeTextCellf(
				"A1",
				"%s",
				xlsx.Value{
					Key: "String",
				},
			),
			expected: testOb{},
		},
	} {
		t.Run(name, func(tt *testing.T) {
			var ob testOb
			elem := reflect.ValueOf(&ob).Elem()
			tc.cell.Encode(elem)
			tc.cell.Decode(elem)
			ob = elem.Interface().(testOb)
			assert.Equal(tt, tc.expected, ob)
		})
	}
}
