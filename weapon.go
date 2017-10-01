package microdungeon

import (
	"github.com/pkg/errors"
)

// Wield a weapon.
func (w *Weapon) Wield(c *Character) error {
    if c == nil {
        return errors.New("c cannot be nil")
    }

	// TODO: check character weapon restrictions.
	//       if this character has a weapon proficiency set a modifier for it.
	if w.Wielded {
		return errors.Errorf("%s is already wielded", w.Name)
	}

	if wielded := c.Wielding; wielded != nil {
		return errors.Errorf("%s is already wielded", wielded.Name)
	}

    c.Wielding = w
	w.Wielded = true

	return nil
}

// Unwield a weapon.
func (w *Weapon) Unwield(c *Character) error {
    if c == nil {
        return errors.New("c cannot be nil")
    }

	if !w.Wielded {
		return errors.Errorf("%s is not wielded", w.Name)
	}

	c.Wielding = nil
	w.Wielded = false

	return nil
}
