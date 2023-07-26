package testdata

import "justanother.org/microdungeon/models"

func Character() models.Character {
	return models.Character{
		Base: models.Base{
			Name: "Mirg",
		},
		PlayerName: "Grim",
		Race: models.Race{
			Base: models.Base{
				Name: "Robot",
			},
			Abilities: models.NewAbilities(map[models.AbilityType]int32{
				models.AbilityTypeStrength:     10,
				models.AbilityTypeIntelligence: 5,
				models.AbilityTypeWisdom:       -2,
				models.AbilityTypeConstitution: 4,
				models.AbilityTypeDexterity:    -3,
			}),
		},
		Class: models.Class{
			Name:   "Peasant",
			HitDie: 12,
			PrimaryAbilities: []models.AbilityType{
				models.AbilityTypeDexterity,
			},
		},
		Alignment: models.Alignment{
			Ethical: -20,
			Moral:   0,
		},
		Score: models.Score{
			Level:      1,
			Experience: 85,
			Abilities: models.NewAbilities(map[models.AbilityType]int32{
				models.AbilityTypeStrength:     -25,
				models.AbilityTypeIntelligence: 10,
				models.AbilityTypeWisdom:       10,
			}),
		},
		Inventory: []models.Object{
			Armour(),
		},
	}
}
