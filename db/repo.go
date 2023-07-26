package db

import (
	"context"

	"justanother.org/microdungeon/models"
)

//go:generate mockgen -build_flags=--mod=mod -destination=./mock/repo.go -package=mock justanother.org/microdungeon/db Repo
type Repo interface {
	StoreRace(ctx context.Context, race models.Race) error
	GetRace(ctx context.Context, id int32) (models.Race, error)
	GetRaceByName(ctx context.Context, name string) (models.Race, error)

	StoreArmour(ctx context.Context, armour models.Armour) error
	GetArmour(ctx context.Context, id int32) (models.Armour, error)
	GetArmourByName(ctx context.Context, name string) (models.Armour, error)

	StoreClass(ctx context.Context, class models.Class) error
	GetClass(ctx context.Context, id int32) (models.Class, error)
	GetClassByName(ctx context.Context, name string) (models.Class, error)

	StoreCharacter(ctx context.Context, char models.Character) error
	GetCharacter(ctx context.Context, id int32) (models.Character, error)
	GetCharacterByName(ctx context.Context, name string) (models.Character, error)
}
