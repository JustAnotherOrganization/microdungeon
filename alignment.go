package microdungeon

const (
	// LawfulGood alignment.
	LawfulGood = "Lawful Good"
	// NeutralGood alignment.
	NeutralGood = "Neutral Good"
	// ChaoticGood alignment.
	ChaoticGood = "Chaotic Good"
	// LawfulNeutral alignment
	LawfulNeutral = "Lawful Neutral"
	// Neutral alignment
	Neutral = "Neutral"
	// ChaoticNeutral alignment
	ChaoticNeutral = "Chaotic Neutral"
	// LawfulEvil alignment
	LawfulEvil = "Lawful Evil"
	// NeutralEvil alignment
	NeutralEvil = "Neutral Evil"
	// ChaoticEvil alignment
	ChaoticEvil = "Chaotic Evil"

	// pad the default values.
	pad = 25

	// Moral thresholds.
	good int32 = 50
	// Neutral exists anywhere between these values.
	evil int32 = -50

	// Ethical thresholds.
	lawful int32 = 50
	// Neutral exists anywhere between these values.
	chaotic = -50
)

// String returns the recognized string name of the given Alignment.
func (a *Alignment) String() string {
	switch {
	case a.Moral >= good && a.Ethical >= lawful:
		return LawfulGood
	case a.Moral >= good && (a.Ethical > chaotic && a.Ethical < lawful):
		return NeutralGood
	case a.Moral >= good && a.Ethical <= chaotic:
		return ChaoticGood
	case (a.Moral > evil && a.Moral < good) && a.Ethical >= lawful:
		return LawfulNeutral
	case (a.Moral > evil && a.Moral < good) && a.Ethical <= chaotic:
		return ChaoticNeutral
	case a.Moral <= evil && a.Ethical >= lawful:
		return LawfulEvil
	case a.Moral <= evil && (a.Ethical > chaotic && a.Ethical < lawful):
		return NeutralEvil
	case a.Moral <= evil && a.Ethical <= chaotic:
		return ChaoticEvil
	//case (a.Moral > evil && a.Moral < good) && (a.Ethical > chaotic && a.Ethical < lawful):
	default:
		return Neutral
	}
}

// LawfulGoodAlignment returns a starting LawfulGood Alignment.
func LawfulGoodAlignment() *Alignment {
	return &Alignment{
		Moral:   good + pad,
		Ethical: lawful + pad,
	}
}

// LawfulNeutralAlignment returns a starting LawfulNeutral Alignment.
func LawfulNeutralAlignment() *Alignment {
	return &Alignment{
		Moral:   good - pad,
		Ethical: lawful + pad,
	}
}

// LawfulEvilAlignment returns a starting LawfulEvil Alignment.
func LawfulEvilAlignment() *Alignment {
	return &Alignment{
		Moral:   evil - pad,
		Ethical: lawful + pad,
	}
}

// NeutralGoodAlignment returns a starting NeutralGood Alignment.
func NeutralGoodAlignment() *Alignment {
	return &Alignment{
		Moral:   good + pad,
		Ethical: lawful - pad,
	}
}

// NeutralAlignment returns a true Neutral Alignment.
func NeutralAlignment() *Alignment {
	return &Alignment{
		Moral:   0,
		Ethical: 0,
	}
}

// NeutralEvilAlignment returns a starting NeutralEvilAlignment.
func NeutralEvilAlignment() *Alignment {
	return &Alignment{
		Moral:   evil - pad,
		Ethical: chaotic + pad,
	}
}

// ChaoticGoodAlignment returns a starting ChaoticGood Alignment.
func ChaoticGoodAlignment() *Alignment {
	return &Alignment{
		Moral:   good + pad,
		Ethical: chaotic - pad,
	}
}

// ChaoticNeutralAlignment returns a starting ChaoticNeutral Alignment.
func ChaoticNeutralAlignment() *Alignment {
	return &Alignment{
		Moral:   evil + pad,
		Ethical: chaotic - pad,
	}
}

// ChaoticEvilAlignment returns a starting ChaoticEvil Alignment.
func ChaoticEvilAlignment() *Alignment {
	return &Alignment{
		Moral:   evil - pad,
		Ethical: chaotic - pad,
	}
}
