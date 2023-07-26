package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"justanother.org/microdungeon/testdata"
)

func TestRepo_Armour(t *testing.T) {
	f := newFixture(t)

	err := f.repo.StoreArmour(context.Background(), testdata.Armour())
	require.NoError(t, err)

	armour, err := f.repo.GetArmourByName(context.Background(), testdata.Armour().Name)
	require.NoError(t, err)

	armour, err = f.repo.GetArmour(context.Background(), armour.Id)
	require.NoError(t, err)

	assert.Equal(t, testdata.Armour().Aliases, armour.Aliases)

	at := armour.Attributes.Get("battery-capacity")
	require.IsType(t, float64(0), at)
	assert.Equal(t, float64(-1), at.(float64))
}
