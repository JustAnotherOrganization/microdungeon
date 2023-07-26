package models

type (
	BaseObject struct {
		Base

		Weight uint32 `json:"weight"`
	}

	Object interface {
		ID(str string) bool
	}
)
