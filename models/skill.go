package models

type Skill struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	DifficultyClass uint32 `json:"DifficultyClass,omitempty"`
	Trained         bool   `json:"trained,omitempty"`
}
