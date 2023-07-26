package xlsx

import (
	"fmt"

	"github.com/xuri/excelize/v2"
	"justanother.org/microdungeon/models"
)

const (
	colWidth = 17.5 // approx 1.5" in libreoffice
)

var sheetMap = map[string]func(char models.Character) interface{}{
	"A1": func(char models.Character) interface{} {
		return fmt.Sprintf("%s (%s, %d)", char.Name, char.PlayerName, char.Id)
	},
	"B1": func(char models.Character) interface{} {
		return fmt.Sprintf("%s / %s", char.Race.Name, char.Class.Name)
	},
	"E1": func(_ models.Character) interface{} {
		return "Alignment Ethical:"
	},
	"F1": func(_ models.Character) interface{} {
		return "Alignment Moral:"
	},

	"A2": func(_ models.Character) interface{} {
		return "AKA"
	},
	"B2": func(char models.Character) interface{} {
		return char.Aliases
	},
	"E2": func(char models.Character) interface{} {
		return char.Alignment.Ethical
	},
	"F2": func(char models.Character) interface{} {
		return char.Alignment.Moral
	},

	"A3": func(char models.Character) interface{} {
		return char.Description
	},

	"A4": func(_ models.Character) interface{} {
		return "Level:"
	},
	"B4": func(char models.Character) interface{} {
		return char.Score.Level
	},
	"C4": func(_ models.Character) interface{} {
		return "Experience:"
	},
	"D4": func(char models.Character) interface{} {
		return char.Score.Experience
	},

	"A6": func(char models.Character) interface{} {
		return fmt.Sprintf("Score (%d):", char.Score.Id)
	},

	"A7": func(_ models.Character) interface{} {
		return models.AbilityTypeStrength.String()
	},
	"B7": func(_ models.Character) interface{} {
		return models.AbilityTypeDexterity.String()
	},
	"C7": func(_ models.Character) interface{} {
		return models.AbilityTypeConstitution.String()
	},
	"D7": func(_ models.Character) interface{} {
		return models.AbilityTypeIntelligence.String()
	},
	"E7": func(_ models.Character) interface{} {
		return models.AbilityTypeWisdom.String()
	},
	"F7": func(_ models.Character) interface{} {
		return models.AbilityTypeCharisma.String()
	},

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

func Character(char models.Character) (*excelize.File, error) {
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
