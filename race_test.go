package microdungeon

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func newTestRace() *Race {
	return &Race{
		Name:  "Zombie Monkeys",
		Genus: "Monkey",
		Description: `The zombie monkeys were created by an evil witch doctor who wanted to destroy
civilization and control the world. They are quick, like lightning and can
rend flesh from bone faster than you can blink.`,
		Score: &Score{
			Set: []*Ability{
				&Ability{Type: AbilityType_STRENGTH, Value: 2},
				&Ability{Type: AbilityType_DEXTERITY, Value: 10},
				&Ability{Type: AbilityType_INTELLIGENCE, Value: -10},
				&Ability{Type: AbilityType_CONSTITION, Value: -2},
				&Ability{Type: AbilityType_WISDOM, Value: -10},
			},
		},
		NightVision:    false,
		KnownLanguages: []string{"common"},
		Undead:         true,
	}
}

var _ = Describe("Race", func() {
	It("works correctly", func() {
		testRace := newTestRace()
		Expect(testRace.Score.Strength()).To(BeNumerically("==", 2))
		Expect(testRace.Score.Dexterity()).To(BeNumerically("==", 10))
		Expect(testRace.Score.Intelligence()).To(BeNumerically("==", -10))
		Expect(testRace.Score.Constitution()).To(BeNumerically("==", -2))
		Expect(testRace.Score.Wisdom()).To(BeNumerically("==", -10))
		Expect(testRace.Score.Charisma()).To(BeNumerically("==", 0))
	})
})
