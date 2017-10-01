package microdungeon

func newTestCharacter() *Character {
	return &Character{
		Name:       "Mirg",
		PlayerName: "Grim",
		Gender:     Character_NEUTER, // This is the default but we're explicit because it's Grim :P
		Race:       newTestRace(),
		Level:      2, // This should not be possible.
	}
}
