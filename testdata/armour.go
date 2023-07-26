package testdata

import "justanother.org/microdungeon/models"

func Armour() models.Armour {
	return models.Armour{
		BaseObject: models.BaseObject{
			Base: models.Base{
				Name: "Cloak of infinite battery power",
				Aliases: []string{
					"cloak",
					"cloak of infinite battery",
					"cloak of battery power",
				},
				Description: "A cloak made of a strange iridescent material. It sparks violently every so often and looks extremely heavy.",
				Attributes: models.NewAttributes(map[models.AttributeName]any{
					"battery-capacity": -1,
				}),
			},
			Weight: 1000,
		},
		AC:   -2,
		Worn: true,
	}
}
