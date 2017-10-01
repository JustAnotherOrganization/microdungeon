package microdungeon

import "github.com/pkg/errors"

func (sc *Score) sum(t AbilityType) (final int32) {
	for _, s := range sc.Set {
		if s.Type == t {
			final += s.Value
		}
	}

	return final
}

// Strength returns the total Strength of the given score set.
func (sc *Score) Strength() int32 {
	return sc.sum(AbilityType_STRENGTH)
}

// Dexterity returns the total Dexterity of the given score set.
func (sc *Score) Dexterity() int32 {
	return sc.sum(AbilityType_DEXTERITY)
}

// Constitution returns the total Constitution of the given score set.
func (sc *Score) Constitution() int32 {
	return sc.sum(AbilityType_CONSTITION)
}

// Intelligence returns the total Intelligence of the given score set.
func (sc *Score) Intelligence() int32 {
	return sc.sum(AbilityType_INTELLIGENCE)
}

// Wisdom returns the total Wisdom of the given score set.
func (sc *Score) Wisdom() int32 {
	return sc.sum(AbilityType_WISDOM)
}

// Charisma returns the total Charisma of the given score set.
func (sc *Score) Charisma() int32 {
	return sc.sum(AbilityType_CHARISMA)
}

// Get the final modifier value from the given modifier.
func (m *Modifier) Get(sc *Score) (int32, error) {
	if sc == nil {
		return 0, errors.New("sc cannot be nil")
	}

	var final int32
	// FIXME: pretty sure this can be done with a logarithm...
	// TODO: ideally this should be part of the dungeon config, not the code.
	switch sc.sum(m.Type) {
	case 1:
		final = -5
	case 2, 3:
		final = -4
	case 4, 5:
		final = -3
	case 6, 7:
		final = -2
	case 8, 9:
		final = -1
	case 10, 11:
		final = 0
	case 12, 13:
		final = 1
	case 14, 15:
		final = 2
	case 16, 17:
		final = 3
	case 18, 19:
		final = 4
	case 20, 21:
		final = 5
	case 22, 23:
		final = 6
	case 24, 25:
		final = 7
	case 26, 27:
		final = 8
	case 28, 29:
		final = 9
	case 30:
		final = 10
	}

	if final > m.Max {
		final = m.Max
	}

	return final, nil
}
