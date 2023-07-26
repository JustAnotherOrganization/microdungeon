package postgres

import (
	"context"
	"strings"

	"go.uber.org/zap"
	"justanother.org/microdungeon/models"
)

func (r Repo) StoreRace(ctx context.Context, race models.Race) error {
	abilitiesID, err := r.storeAbilities(ctx, race.Abilities)
	if err != nil {
		return err
	}

	const sql = `
	INSERT INTO microdungeon.races (
		name, aliases, description, abilities, attributes
	) VALUES (
		$1, $2, $3, $4, $5
	) ON CONFLICT (name) DO UPDATE
	SET
		aliases = $2,
		description = $3,
		abilities = $4,
		attributes = $5
	;`

	_, err = r.pool.Exec(ctx, sql,
		race.Name,
		race.Aliases,
		race.Description,
		abilitiesID,
		race.Attributes.AsMap())
	return err
}

func (r Repo) GetRace(ctx context.Context, id int32) (models.Race, error) {
	return r.getRace(ctx, byIdClause, id)
}

func (r Repo) GetRaceByName(ctx context.Context, name string) (models.Race, error) {
	return r.getRace(ctx, byNameClause, name)
}

func (r Repo) getRace(ctx context.Context, clause string, params ...any) (models.Race, error) {
	const _sql = `SELECT id, name, aliases, description, abilities, attributes FROM microdungeon.races`

	var b strings.Builder
	b.WriteString(_sql)
	b.WriteString("\n")
	b.WriteString(clause)
	b.WriteString("\n;")

	sql := b.String()

	r.log.
		Debug("query", zap.String("sql", sql))

	var (
		race        models.Race
		abilitiesID int32
	)

	err := r.pool.QueryRow(ctx, sql, params...).Scan(
		&race.Id,
		&race.Name,
		&race.Aliases,
		&race.Description,
		&abilitiesID,
		&race.Attributes,
	)
	if err != nil {
		return race, err
	}

	race.Abilities, err = r.getAbilities(ctx, abilitiesID)
	if err != nil {
		return race, err
	}

	return race, nil
}

/*
func (r Repo) ListRaces(ctx context.Context, filterBy ...models.AttributeName) ([]models.Race, error) {
	const sql = `
	SELECT name, genus, description, strength, dexterity, constitution,
	intelligence, wisdom, charisma, attributes FROM microdungeon.races
	WHERE
`
}
*/
