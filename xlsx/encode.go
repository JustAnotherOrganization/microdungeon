package xlsx

import (
	"reflect"
)

func (s *SheetMap) encode(ob interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()

	to := reflect.TypeOf(ob)

	var elem reflect.Value
	if to.Kind() == reflect.Ptr {
		elem = reflect.ValueOf(ob).Elem()
	} else {
		elem = reflect.ValueOf(ob)
	}

	for idx, cell := range s.cells {
		cell.Encode(elem)
		s.cells[idx] = cell
	}
}
