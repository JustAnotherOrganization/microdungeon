package models

const (
	LawfulGood     = "Lawful Good"
	NeutralGood    = "Neutral Good"
	ChaoticGood    = "Chaotic Good"
	LawfulNeutral  = "Lawful Neutral"
	Neutral        = "Neutral"
	ChaoticNeutral = "Chaotic Neutral"
	LawfulEvil     = "Lawful Evil"
	NeutralEvil    = "Neutral Evil"
	ChaoticEvil    = "Chaotic Evil"

	// Moral thresholds.
	good int32 = 50
	// Neutral exists anywhere between these values.
	evil int32 = -50

	// Ethical thresholds.
	lawful int32 = 50
	// Neutral exists anywhere between these values.
	chaotic = -50
)

type Alignment struct {
	Ethical int32 `json:"ethical"`
	Moral   int32 `json:"moral"`
}

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
