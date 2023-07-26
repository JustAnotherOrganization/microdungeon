package models

var _ = Describe("Alignment", func() {
	It("works correctly", func() {
		lga := LawfulGoodAlignment()
		lna := LawfulNeutralAlignment()
		lea := LawfulEvilAlignment()
		nga := NeutralGoodAlignment()
		na := NeutralAlignment()
		nea := NeutralEvilAlignment()
		cga := ChaoticGoodAlignment()
		cna := ChaoticNeutralAlignment()
		cea := ChaoticEvilAlignment()

		testCases := []struct {
			moral    int32
			ethical  int32
			expected string
		}{
			{50, 50, LawfulGood},
			{50, 49, NeutralGood},
			{50, -49, NeutralGood},
			{50, -50, ChaoticGood},
			{49, -49, Neutral},
			{-49, 49, Neutral},
			{-49, -49, Neutral},
			{0, 0, Neutral},
			{-50, 50, LawfulEvil},
			{-50, 49, NeutralEvil},
			{-50, -49, NeutralEvil},
			{-50, -50, ChaoticEvil},
			{lga.Moral, lga.Ethical, LawfulGood},
			{lna.Moral, lna.Ethical, LawfulNeutral},
			{lea.Moral, lea.Ethical, LawfulEvil},
			{nga.Moral, nga.Ethical, NeutralGood},
			{na.Moral, na.Ethical, Neutral},
			{nea.Moral, nea.Ethical, NeutralEvil},
			{cga.Moral, cga.Ethical, ChaoticGood},
			{cna.Moral, cna.Ethical, ChaoticNeutral},
			{cea.Moral, cea.Ethical, ChaoticEvil},
		}

		for _, tc := range testCases {
			Expect((&Alignment{
				Moral:   tc.moral,
				Ethical: tc.ethical,
			}).String()).To(Equal(tc.expected))
		}
	})
})
