package models

//go:generate enumer -type=DamageType
type DamageType int32

const (
	DamageTypeUnknown DamageType = iota
	DamageTypeBludgeoning
	DamageTypePiercing
	DamageTypeSlashing
)
