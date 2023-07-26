package models

type Character struct {
	Base

	PlayerName string    `json:"player-name"`
	Score      Score     `json:"score"`
	Alignment  Alignment `json:"alignment"`
	Inventory  []Object  `json:"inventory,omitempty"`

	Race  Race  `json:"race"`
	Class Class `json:"class"`
}
