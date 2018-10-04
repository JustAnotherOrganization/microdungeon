package microdungeon

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func newTestCharacter() *Character {
	return &Character{
		Name:       "Mirg",
		PlayerName: "Grim",
		Gender:     Character_NEUTER, // This is the default but we're explicit because it's Grim :P
		Race:       newTestRace(),
		Level:      2, // This should not be possible.
		Class:      newTestClass(),
	}
}

var _ = Describe("Character", func() {
	It("", func() {
		a := newTestArmour(Armour_MAIN, WeightClass_HEAVY)
		err := a.Wear(nil)
		Expect(err).ToNot(BeNil())
	})
	It("can not weild a restricted type", func() {
		c := newTestCharacter()
		weapon := &Weapon{
			Type: Weapon_CANNON,
		}
		err := c.Wield(weapon)
		Expect(err).To(Not(BeNil()))
	})
})
