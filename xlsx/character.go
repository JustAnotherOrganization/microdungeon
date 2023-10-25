package xlsx

import (
	"errors"

	"github.com/xuri/excelize/v2"
	"justanother.org/microdungeon/models"
)

func ExcelFromCharacter(char models.Character) (*excelize.File, error) {
	m := NewSheetMap(
		NewCompositeTextCellf("A1", "%s (%s, %d)",
			Value{Key: "Name"},
			Value{Key: "PlayerName"},
			Value{Key: "Id"}),
		NewCompositeTextCellf("B1", "%s / %s",
			Value{Key: "Race.Name"},
			Value{Key: "Class.Name"}),
		NewTextCell("E1", "", "Alignment Ethical:"),
		NewTextCell("F1", "", "Alignment Moral:"),
		NewTextCell("A2", "", "AKA"),
		// FIXME: handle []string and other slices (instead of relying on fmt's %v)
		NewTextCell("B2", "Aliases", nil),
		NewTextCell("E2", "Alignment.Ethical", nil),
		NewTextCell("F2", "Alignment.Moral", nil),
		NewTextCell("A3", "Description", nil),
		NewTextCell("A4", "", "Level:"),
		NewTextCell("B4", "Score.Level", nil),
		NewTextCell("C4", "", "Experience:"),
		NewTextCell("D4", "Score.Experience", nil),
		NewCompositeTextCellf("A6", "Score (%d):",
			Value{Key: "Score.Id"}),
		NewTextCell("A7", "", models.AbilityTypeStrength.String()),
		NewTextCell("B7", "", models.AbilityTypeDexterity.String()),
		NewTextCell("C7", "", models.AbilityTypeConstitution.String()),
		NewTextCell("D7", "", models.AbilityTypeIntelligence.String()),
		NewTextCell("E7", "", models.AbilityTypeWisdom.String()),
		NewTextCell("F7", "", models.AbilityTypeCharisma.String()),
		// FIXME: need to be able to support function calls in some manner...
		// 		eq, char.Score.Abilities.Get(models.AbilityTypeStrength)
	)
	m = m
	return nil, errors.New("not yet implemented")
}
