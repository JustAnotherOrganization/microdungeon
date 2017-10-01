package microdungeon

// Do the skill.
func (sk *Skill) Do(roll int32, char *Character) (int32, bool, error) {
	final := roll + int32(char.Level/2)
	for _, m := range sk.Modifiers {
		v, err := m.Get(char.Score)
		if err != nil {
			return final, false, err
		}

		final += v
	}

	if sk.Trained {
		final += 5
	}

	if final > int32(sk.DifficultyClass) {
		return final, true, nil
	}

	return final, false, nil
}
