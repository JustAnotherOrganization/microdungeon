package xlsx

//func (s *SheetMap) decode(ob interface{}) error {
//	s.mux.Lock()
//	defer s.mux.Unlock()
//
//	to := reflect.TypeOf(ob)
//
//	if to.Kind() != reflect.Ptr || to.Elem().Kind() != reflect.Struct {
//		return ErrObjectNotPtr
//	}
//
//	for _, cell := range s.cells {
//		cell.Decode(ob)
//	}
//
//	return nil
//}
