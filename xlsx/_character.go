package xlsx

import (
	"github.com/xuri/excelize/v2"
	"justanother.org/microdungeon/models"
)

const (
	colWidth = 17.5 // approx 1.5" in libreoffice
)

type cellMap struct {
	ToVal   func(char models.Character) interface{}
	FromVal func(char *models.Character, val string) error
}

var sheetMap = map[string]cellMap{

	"A8": func(char models.Character) interface{} {
		return char.Score.Abilities.Get(models.AbilityTypeStrength)
	},
	"B8": func(char models.Character) interface{} {
		return char.Score.Abilities.Get(models.AbilityTypeDexterity)
	},
	"C8": func(char models.Character) interface{} {
		return char.Score.Abilities.Get(models.AbilityTypeConstitution)
	},
	"D8": func(char models.Character) interface{} {
		return char.Score.Abilities.Get(models.AbilityTypeIntelligence)
	},
	"E8": func(char models.Character) interface{} {
		return char.Score.Abilities.Get(models.AbilityTypeWisdom)
	},
	"F8": func(char models.Character) interface{} {
		return char.Score.Abilities.Get(models.AbilityTypeCharisma)
	},
}

func ExcelFromCharacter(char models.Character) (*excelize.File, error) {
	f := excelize.NewFile()

	sheet, err := f.NewSheet(char.Name)
	if err != nil {
		return nil, err
	}

	if err = f.DeleteSheet("Sheet1"); err != nil {
		return nil, err
	}

	f.SetActiveSheet(sheet)

	if err = f.SetColWidth(char.Name, "A", "H", colWidth); err != nil {
		return nil, err
	}

	for col, val := range sheetMap {
		if err = f.SetCellValue(char.Name, col, val(char)); err != nil {
			return nil, err
		}
	}

	// attributes
	// inventory
	return f, nil
}

func CharacterFromExcel(name string, f *excelize.File) (models.Character, error) {
	f.GetCellValue(name)

}
