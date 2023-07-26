package models

type Class struct {
	Id               int32         `json:"id"`
	Name             string        `json:"name"`
	HitDie           uint32        `json:"hit-die"`
	PrimaryAbilities []AbilityType `json:"primary-abilities"`
	Saves            []AbilityType `json:"saves"`
}
