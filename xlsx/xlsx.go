package xlsx

import (
	"errors"
	"sync"

	"github.com/xuri/excelize/v2"
)

type (
	SheetMap struct {
		mux   sync.Mutex
		cells []Cell
	}
)

var ErrObjectNotPtr = errors.New("ob must be a pointer to a struct")

func NewSheetMap(cells ...Cell) *SheetMap {
	return &SheetMap{
		cells: cells,
	}
}

func ExcelFromOb(name string, sheetMap *SheetMap, ob interface{}, opt ...Option) (*excelize.File, error) {
	f := excelize.NewFile()

	sheet, err := f.NewSheet(name)
	if err != nil {
		return nil, err
	}

	f.SetActiveSheet(sheet)

	if err = f.DeleteSheet("Sheet1"); err != nil {
		return nil, err
	}

	for _, o := range opt {
		if err = o.Apply(name, sheetMap, f); err != nil {
			return nil, err
		}
	}

	sheetMap.encode(ob)

	sheetMap.mux.Lock()
	defer sheetMap.mux.Unlock()

	for _, cell := range sheetMap.cells {
		if err = f.SetCellValue(name, cell.Location(), cell.Interface()); err != nil {
			return nil, err
		}
	}

	return f, nil
}
