package microdungeon

func newTestClass() *Class {
	return &Class{
		Name:   "Cleric",
		HitDie: 42,
		PrimaryAbilities: []AbilityType{
			AbilityType_INTELLIGENCE,
		},
		WeaponTypeRestrictions: []Weapon_Type{
			Weapon_CANNON,
		},
	}
}
