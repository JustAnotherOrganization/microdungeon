package models

type WeaponType int32

const (
	WeaponUnknown WeaponType = iota
	WeaponDagger
	WeaponShortBlade
	WeaponLongBlade
	WeaponHammer
	WeaponAxe
	WeaponPoleArm
	WeaponBow
	WeaponCrossbow
	WeaponBlackPowerPistol
	WeaponBlackPowderRifle
	WeaponPistol
	WeaponRifle
	WeaponCannon
)

// Ammunition implements ammunition for projectile weapons.
type Ammunition struct {
	BaseObject

	DamageType DamageType `json:"damage_type,omitempty"`
	WC         int32      `json:"wc,omitempty"`
}

// Weapon is the basis for all weapons.
type Weapon struct {
	BaseObject

	Type       WeaponType `json:"type,omitempty"`
	DamageType DamageType `json:"damage_type,omitempty"`
	WC         int32      `json:"wc,omitempty"`
	Weight     uint32     `json:"weight,omitempty"`
}

/*
func (w *Weapon) Wield(c *Character) error {
	if c == nil {
		return errors.New("c cannot be nil")
	}

	// TODO: check character weapon restrictions.
	//       if this character has a weapon proficiency set a modifier for it.
	if w.Wielded {
		return fmt.Errorf("%s is already wielded", w.Name)
	}

	if wielded := c.Wielding; wielded != nil {
		return fmt.Errorf("%s is already wielded", wielded.Name)
	}

	c.Wielding = w
	w.Wielded = true

	return nil
}

func (w *Weapon) Unwield(c *Character) error {
	if c == nil {
		return errors.New("c cannot be nil")
	}

	if !w.Wielded {
		return fmt.Errorf("%s is not wielded", w.Name)
	}

	c.Wielding = nil
	w.Wielded = false

	return nil
}
*/
