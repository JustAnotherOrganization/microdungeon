package models

type Race struct {
	Base

	Abilities *Abilities `json:"abilities,omitempty"`
}
