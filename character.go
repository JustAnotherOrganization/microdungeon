package microdungeon

import (
	"fmt"
)

func (c *Character) Wield(weapon *Weapon) error {
	if c.IsWeaponRestricted(weapon) {
		return fmt.Errorf("%s is not allowed for this character class", weapon.Type.String())
	}

	if weapon.Wielded {
		return fmt.Errorf("%s is already wielded", weapon.Name)
	}

	if c.Wielding != nil {
		return fmt.Errorf("%s is already wielded", c.Wielding.Name)
	}

	c.Wielding = weapon
	weapon.Wielded = true

	return nil
}

func (c *Character) IsWeaponRestricted(weapon *Weapon) bool {
	for _, restricted := range c.Class.WeaponTypeRestrictions {
		if weapon.Type == restricted {
			return true
		}
	}
	return false
}
