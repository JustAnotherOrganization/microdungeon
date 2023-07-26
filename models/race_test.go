package models_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"justanother.org/microdungeon/models"
	"justanother.org/microdungeon/testdata"
)

func TestRace(t *testing.T) {
	tr := testdata.Race()

	byt, err := json.Marshal(tr)
	require.NoError(t, err)

	var ntr models.Race
	err = json.Unmarshal(byt, &ntr)
	require.NoError(t, err)

	assert.Equal(t, int32(2), ntr.Abilities.Get(models.AbilityTypeStrength))
	assert.Equal(t, int32(10), ntr.Abilities.Get(models.AbilityTypeDexterity))
	assert.Equal(t, int32(-10), ntr.Abilities.Get(models.AbilityTypeIntelligence))
	assert.Equal(t, int32(-2), ntr.Abilities.Get(models.AbilityTypeConstitution))
	assert.Equal(t, int32(-10), ntr.Abilities.Get(models.AbilityTypeWisdom))
	assert.Equal(t, int32(0), ntr.Abilities.Get(models.AbilityTypeCharisma))

	/*
		at, ok := tr.Attributes.ForName(models.Genus)
		assert.True(t, ok)
		assert.IsType(t, "", at.Value())
		assert.Equal(t, "Monkey", at.Value().(string))

		_, ok = tr.Attributes.ForName(models.Playable)
		assert.False(t, ok)

		at, ok = tr.Attributes.ForName(models.NightVision)
		assert.True(t, ok)
		assert.IsType(t, true, at.Value())
		assert.False(t, at.Value().(bool))

		at, ok = tr.Attributes.ForName(models.KnownLanguage)
		assert.True(t, ok)
		assert.IsType(t, []string{}, at.Value())
		sl := at.Value().([]string)
		assert.Len(t, sl, 1)
		assert.Equal(t, "monkey", sl[0])

		at, ok = tr.Attributes.ForName(models.Undead)
		assert.IsType(t, true, at.Value())
		assert.True(t, at.Value().(bool))
	*/
}
