package xlsx

import "github.com/xuri/excelize/v2"

type (
	Option interface {
		Apply(name string, sheetMap *SheetMap, f *excelize.File) error
	}

	ColumnWidthOption struct {
		StartCol string
		EndCol   string
		Width    float64
	}
)

func (o ColumnWidthOption) Apply(name string, _ *SheetMap, f *excelize.File) error {
	return f.SetColWidth(name, o.StartCol, o.EndCol, o.Width)
}
