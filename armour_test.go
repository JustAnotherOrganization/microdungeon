package microdungeon

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func newTestArmour(_type Armour_Type, wc WeightClass) *Armour {
	return &Armour{
		Name:        "A piece of test armour",
		Aliases:     []string{"test armour", "piece"},
		Description: `A piece of test armour, this is it's long description`,
		WeightClass: wc,
		Type:        _type,
	}
}

var _ = Describe("Armour", func() {
	It("Errors if the character provided is nil", func() {
		a := newTestArmour(Armour_MAIN, WeightClass_HEAVY)
		err := a.Wear(nil)
		Expect(err).ToNot(BeNil())
	})

	Context("Enforce the number of worn armour types", func() {
		char := newTestCharacter()

		It("ensures only one piece of headgear can be worn", func() {
			a := newTestArmour(Armour_HEADGEAR, WeightClass_LIGHT)
			err := a.Wear(char)
			Expect(err).To(BeNil())

			a = newTestArmour(Armour_HEADGEAR, WeightClass_LIGHT)
			err = a.Wear(char)
			Expect(err).ToNot(BeNil())
		})

		It("ensures only one piece of main armour can be worn", func() {
			a := newTestArmour(Armour_MAIN, WeightClass_LIGHT)
			err := a.Wear(char)
			Expect(err).To(BeNil())

			a = newTestArmour(Armour_MAIN, WeightClass_LIGHT)
			err = a.Wear(char)
			Expect(err).ToNot(BeNil())
		})

		It("ensures only one pair of gloves can be worn", func() {
			a := newTestArmour(Armour_GLOVES, WeightClass_LIGHT)
			err := a.Wear(char)
			Expect(err).To(BeNil())

			a = newTestArmour(Armour_GLOVES, WeightClass_LIGHT)
			err = a.Wear(char)
			Expect(err).ToNot(BeNil())
		})

		It("ensures only one pair of boots can be worn", func() {
			a := newTestArmour(Armour_BOOTS, WeightClass_LIGHT)
			err := a.Wear(char)
			Expect(err).To(BeNil())

			a = newTestArmour(Armour_BOOTS, WeightClass_LIGHT)
			err = a.Wear(char)
			Expect(err).ToNot(BeNil())
		})
		// FIXME:
		/*
			It("ensures only two rings can be worn", func() {
				a := newTestArmour(Armour_RING, WeightClass_LIGHT)
				err := a.Wear(char)
				Expect(err).To(BeNil())

				a = newTestArmour(Armour_RING, WeightClass_LIGHT)
				err = a.Wear(char)
				Expect(err).To(BeNil())

				a = newTestArmour(Armour_RING, WeightClass_LIGHT)
				err = a.Wear(char)
				Expect(err).ToNot(BeNil())
			})
		*/
		It("ensures only one belt can be worn", func() {
			a := newTestArmour(Armour_BELT, WeightClass_LIGHT)
			err := a.Wear(char)
			Expect(err).To(BeNil())

			a = newTestArmour(Armour_BELT, WeightClass_LIGHT)
			err = a.Wear(char)
			Expect(err).ToNot(BeNil())
		})

		It("ensures only one cloak can be worn", func() {
			a := newTestArmour(Armour_CLOAK, WeightClass_LIGHT)
			err := a.Wear(char)
			Expect(err).To(BeNil())

			a = newTestArmour(Armour_CLOAK, WeightClass_LIGHT)
			err = a.Wear(char)
			Expect(err).ToNot(BeNil())
		})

		It("ensures only one shield can be worn", func() {
			a := newTestArmour(Armour_SHIELD, WeightClass_LIGHT)
			err := a.Wear(char)
			Expect(err).To(BeNil())

			a = newTestArmour(Armour_SHIELD, WeightClass_LIGHT)
			err = a.Wear(char)
			Expect(err).ToNot(BeNil())
		})

		It("ensures only one pair of glasses can be worn", func() {
			a := newTestArmour(Armour_GLASSES, WeightClass_LIGHT)
			err := a.Wear(char)
			Expect(err).To(BeNil())

			a = newTestArmour(Armour_GLASSES, WeightClass_LIGHT)
			err = a.Wear(char)
			Expect(err).ToNot(BeNil())
		})
		// FIXME:
		/*
			It("ensures only three pieces of unknown gear can be worn", func() {
				a := newTestArmour(Armour_UNKNOWN, WeightClass_LIGHT)
				err := a.Wear(char)
				Expect(err).To(BeNil())

				a = newTestArmour(Armour_UNKNOWN, WeightClass_LIGHT)
				err = a.Wear(char)
				Expect(err).To(BeNil())

				a = newTestArmour(Armour_UNKNOWN, WeightClass_LIGHT)
				err = a.Wear(char)
				Expect(err).To(BeNil())

				a = newTestArmour(Armour_UNKNOWN, WeightClass_LIGHT)
				err = a.Wear(char)
				Expect(err).ToNot(BeNil())
			})
		*/
	})
})
