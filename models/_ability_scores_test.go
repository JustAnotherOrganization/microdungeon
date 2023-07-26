package models

var _ = Describe("Modifier", func() {
	It("works correctly", func() {
		// TODO: improve this when replacing the switch statement.
		m := Modifier{
			Type: AbilityTypeStrength,
		}
		sc := Score{
			Set: []*Ability{
				&Ability{
					Type:  AbilityTypeStrength,
					Value: 10,
				},
			},
		}

		n, err := m.Get(&sc)
		Expect(err).To(BeNil())
		Expect(n).To(BeNumerically("==", 0))
	})
})
