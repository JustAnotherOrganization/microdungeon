package models

import (
	"encoding/json"
	"sync"
)

type AttributeName string

const (
	Genus          AttributeName = "genus"
	Playable       AttributeName = `playable`
	KnownLanguages AttributeName = `known-languages`
	NightVision    AttributeName = `night-vision`
	Undead         AttributeName = `undead`
)

type Attributes struct {
	mux sync.RWMutex
	m   map[AttributeName]any
}

func NewAttributes(m map[AttributeName]any) *Attributes {
	a := new(Attributes)
	a.FromMap(m)
	return a
}

func (a *Attributes) Set(name AttributeName, value any) {
	if a == nil {
		return
	}

	a.mux.Lock()
	defer a.mux.Unlock()

	a.m[name] = value
}

func (a *Attributes) Get(name AttributeName) any {
	if a == nil {
		return nil
	}

	a.mux.RLock()
	defer a.mux.RUnlock()

	return a.m[name]
}

func (a *Attributes) AsMap() map[AttributeName]any {
	if a == nil {
		return nil
	}

	a.mux.RLock()
	defer a.mux.RUnlock()

	nm := make(map[AttributeName]any, len(a.m))
	for k, v := range a.m {
		nm[k] = v
	}

	return nm
}

func (a *Attributes) FromMap(m map[AttributeName]any) {
	if a == nil {
		return
	}

	a.mux.Lock()
	defer a.mux.Unlock()

	a.m = make(map[AttributeName]any, len(m))
	for k, v := range m {
		a.m[k] = v
	}
}

func (a *Attributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.AsMap())
}

func (a *Attributes) UnmarshalJSON(data []byte) error {
	m := make(map[AttributeName]any)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	a.FromMap(m)
	return nil
}
