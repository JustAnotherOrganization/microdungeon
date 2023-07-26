package models

import "strconv"

type Base struct {
	Id          int32       `json:"id"`
	Name        string      `json:"name"`
	Aliases     []string    `json:"aliases,omitempty"`
	Description string      `json:"description,omitempty"`
	Attributes  *Attributes `json:"attributes,omitempty"`
}

const base10 = 10

func (b Base) ID(str string) bool {
	if str == strconv.FormatInt(int64(b.Id), base10) ||
		str == b.Name {
		return true
	}

	for _, a := range b.Aliases {
		if str == a {
			return true
		}
	}

	return false
}
