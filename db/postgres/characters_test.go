package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"justanother.org/microdungeon/models"
	"justanother.org/microdungeon/testdata"
)

func TestRepo_Character(t *testing.T) {
	f := newFixture(t)

	err := f.repo.StoreRace(context.Background(), testdata.Character().Race)
	require.NoError(t, err)

	err = f.repo.StoreClass(context.Background(), testdata.Character().Class)
	require.NoError(t, err)

	err = f.repo.StoreCharacter(context.Background(), testdata.Character())
	require.NoError(t, err)

	char, err := f.repo.GetCharacterByName(context.Background(), testdata.Character().Name)
	require.NoError(t, err)

	char, err = f.repo.GetCharacter(context.Background(), char.Id)
	require.NoError(t, err)

	assert.Equal(t, "Robot", char.Race.Name)
	assert.Equal(t, "Peasant", char.Class.Name)
	require.Len(t, char.Inventory, 1)
	require.IsType(t, models.Armour{}, char.Inventory[0])
	armour := char.Inventory[0].(models.Armour)
	assert.Equal(t, testdata.Armour().Name, armour.Name)
	assert.True(t, armour.Worn)
}
