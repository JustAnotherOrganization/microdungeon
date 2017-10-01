package microdungeon

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Modifier", func() {
	It("works correctly", func() {
		// TODO: improve this when replacing the switch statement.
		m := Modifier{
			Type: AbilityType_STRENGTH,
		}
		sc := Score{
			Set: []*Ability{
				&Ability{
					Type:  AbilityType_STRENGTH,
					Value: 10,
				},
			},
		}

		n, err := m.Get(&sc)
		Expect(err).To(BeNil())
		Expect(n).To(BeNumerically("==", 0))
	})
})
