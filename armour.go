package microdungeon

import "github.com/pkg/errors"

var (
	// MaxRings is the max number of rings a character can wear.
	// TODO: make this dependent on race (we may want to support a race with
	// 3 hands for example).
	MaxRings = 2
	// MaxUnknown is the max number of unknown items a character can wear.
	MaxUnknown = 3
)

func (a *Armour) checkWear(c *Character) error {
	if c == nil {
		return errors.New("c cannot be nil")
	}

	var (
		rings   = make([]*Armour, 0, MaxRings)
		unknown = make([]*Armour, 0, MaxUnknown)
	)

	for _, _a := range c.Armour {
		if _a.Type == a.Type {
			switch a.Type {
			case Armour_HEADGEAR, Armour_MAIN, Armour_GLOVES, Armour_BOOTS,
				Armour_BELT, Armour_CLOAK, Armour_SHIELD, Armour_GLASSES:
				return errors.Errorf("%s is already worn", _a.Name)
			case Armour_RING:
				rings = append(rings, _a)

				if len(rings) >= MaxRings {
					return errors.Errorf("%s and %s are already worn", rings[0].Name, rings[1].Name)
				}
			default:
				unknown = append(unknown, _a)

				if len(unknown) >= MaxUnknown {
					return errors.Errorf("%s, %s and %s are already worn", unknown[0].Name, unknown[1].Name, unknown[2].Name)
				}
			}
		}
	}

	return nil
}

// Wear a piece of armour.
func (a *Armour) Wear(c *Character) error {
	if c == nil {
		return errors.New("c cannot be nil")
	}

	// TODO: check character armour restrictions.
	//       if this character has a armour proficiency set a modifier for it.
	if a.Worn {
		return errors.Errorf("%s is already worn", a.Name)
	}

	if err := a.checkWear(c); err != nil {
		return err
	}

	c.Armour = append(c.Armour, a)
	a.Worn = true

	return nil
}
