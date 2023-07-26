package testdata

import (
	"justanother.org/microdungeon/models"
)

func Race() models.Race {
	return models.Race{
		Base: models.Base{
			Name: "Zombie Monkeys",
			Description: `The zombie monkeys were created by an evil witch doctor who wanted to destroy
civilization and control the world. They are quick, like lightning and can
rend flesh from bone faster than you can blink.`,
			Attributes: models.NewAttributes(map[models.AttributeName]any{
				models.Genus:          "Monkey",
				models.NightVision:    false,
				models.Undead:         true,
				models.KnownLanguages: []string{"monkey"},
			}),
		},
		Abilities: models.NewAbilities(map[models.AbilityType]int32{
			models.AbilityTypeStrength:     2,
			models.AbilityTypeDexterity:    10,
			models.AbilityTypeIntelligence: -10,
			models.AbilityTypeConstitution: -2,
			models.AbilityTypeWisdom:       -10,
		}),
	}
}
