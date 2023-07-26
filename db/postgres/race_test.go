package postgres_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"justanother.org/microdungeon/models"
	"justanother.org/microdungeon/testdata"
)

func TestRepo_Races(t *testing.T) {
	f := newFixture(t)

	err := f.repo.StoreRace(context.Background(), testdata.Race())
	require.NoError(t, err)

	race, err := f.repo.GetRaceByName(context.Background(), testdata.Race().Name)
	require.NoError(t, err)

	race, err = f.repo.GetRace(context.Background(), race.Id)
	require.NoError(t, err)

	at := race.Attributes.Get(models.Genus)
	require.NotNil(t, at)
	assert.IsType(t, "", at)
	assert.Equal(t, "Monkey", at.(string))
}
