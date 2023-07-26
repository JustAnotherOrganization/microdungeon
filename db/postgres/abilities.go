package postgres

import (
	"context"
	"strings"

	"go.uber.org/zap"
	"justanother.org/microdungeon/models"
)

func (r Repo) storeAbilities(ctx context.Context, abils *models.Abilities) (int32, error) {
	if abils.Id == 0 {
		return r._storeAbilities(ctx, "", abils)
	}

	return r._storeAbilities(ctx, "WHERE id = $7", abils)
}

func (r Repo) _storeAbilities(ctx context.Context, clause string, abils *models.Abilities) (int32, error) {
	const _sql = `
	INSERT INTO microdungeon.abilities (
		strength, dexterity, constitution,
		intelligence, wisdom, charisma 
	) VALUES (
		$1, $2, $3,
		$4, $5, $6
	)`

	params := make([]any, 0, 7)
	params = append(params, abils.Get(models.AbilityTypeStrength))
	params = append(params, abils.Get(models.AbilityTypeDexterity))
	params = append(params, abils.Get(models.AbilityTypeConstitution))
	params = append(params, abils.Get(models.AbilityTypeIntelligence))
	params = append(params, abils.Get(models.AbilityTypeWisdom))
	params = append(params, abils.Get(models.AbilityTypeCharisma))

	var b strings.Builder
	b.WriteString(_sql)
	b.WriteString("\n")
	if clause != "" {
		b.WriteString(clause)
		b.WriteString("\n")
		params = append(params, abils.Id)
	}
	b.WriteString("RETURNING ID\n;")

	sql := b.String()

	r.log.
		Debug("query", zap.String("sql", sql))

	var id int32
	err := r.pool.QueryRow(ctx, sql, params...).Scan(&id)
	return id, err
}

func (r Repo) getAbilities(ctx context.Context, id int32) (*models.Abilities, error) {
	const sql = `
	SELECT
		strength, dexterity, constitution,
		intelligence, wisdom, charisma
	FROM microdungeon.abilities
	WHERE id = $1
	;`

	var str, dex, con, intel, wis, char int32
	err := r.pool.QueryRow(ctx, sql, id).Scan(
		&str,
		&dex,
		&con,
		&intel,
		&wis,
		&char,
	)
	if err != nil {
		return nil, err
	}

	abils := new(models.Abilities)
	abils.Set(models.AbilityTypeStrength, str)
	abils.Set(models.AbilityTypeDexterity, dex)
	abils.Set(models.AbilityTypeConstitution, con)
	abils.Set(models.AbilityTypeIntelligence, intel)
	abils.Set(models.AbilityTypeWisdom, wis)
	abils.Set(models.AbilityTypeCharisma, char)

	return abils, nil
}
