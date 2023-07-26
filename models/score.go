package models

import (
	"encoding/json"
	"sync"
)

type (
	Score struct {
		Id         int32      `json:"id"`
		Experience uint32     `json:"experience"`
		Level      uint32     `json:"level"`
		Abilities  *Abilities `json:"abilities"`
	}

	Abilities struct {
		Id  int32 `json:"id"`
		mux sync.RWMutex
		m   map[AbilityType]int32
	}
)

func NewAbilities(m map[AbilityType]int32) *Abilities {
	a := new(Abilities)
	a.FromMap(m)
	return a
}

func (a *Abilities) Set(t AbilityType, val int32) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.m == nil {
		a.m = make(map[AbilityType]int32)
	}

	a.m[t] = val
}

func (a *Abilities) Get(t AbilityType) int32 {
	a.mux.RLock()
	defer a.mux.RUnlock()

	if a.m == nil {
		return 0
	}

	return a.m[t]
}

func (a *Abilities) FromMap(m map[AbilityType]int32) {
	if a == nil {
		return
	}

	a.mux.Lock()
	defer a.mux.Unlock()

	a.m = make(map[AbilityType]int32)
	for k, v := range m {
		a.m[k] = v
	}
}

func (a *Abilities) MarshalJSON() ([]byte, error) {
	a.mux.RLock()
	defer a.mux.RUnlock()

	return json.Marshal(a.m)
}

func (a *Abilities) UnmarshalJSON(data []byte) error {
	a.mux.Lock()
	defer a.mux.Unlock()

	a.m = make(map[AbilityType]int32)
	return json.Unmarshal(data, &a.m)
}
