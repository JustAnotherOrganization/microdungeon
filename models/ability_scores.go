package models

import "encoding/json"

//go:generate enumer -type=AbilityType
type AbilityType int32

const (
	AbilityTypeUnknown AbilityType = iota
	AbilityTypeStrength
	AbilityTypeDexterity
	AbilityTypeConstitution
	AbilityTypeIntelligence
	AbilityTypeWisdom
	AbilityTypeCharisma
)

func (at AbilityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(at.String())
}

type Ability struct {
	Type  AbilityType `json:"type"`
	Value int32       `json:"value"`
}

func (a *Ability) UnmarshalJSON(data []byte) error {
	type iface struct {
		Type  string `json:"type"`
		Value int32  `json:"value"`
	}

	var val iface
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	at, err := AbilityTypeString(val.Type)
	if err != nil {
		return err
	}

	a.Type = at
	a.Value = val.Value
	return nil
}
