package postgres

import (
	"context"
	"strings"

	"go.uber.org/zap"
	"justanother.org/microdungeon/models"
)

func (r Repo) StoreArmour(ctx context.Context, armour models.Armour) error {
	_, err := r.storeArmour(ctx, armour)
	return err
}

func (r Repo) storeArmour(ctx context.Context, armour models.Armour) (int32, error) {
	const sql = `
	INSERT INTO microdungeon.armour (
		name, aliases, description, weight,
		type, ac, worn, attributes
	) VALUES (
		$1, $2, $3, $4,
		$5, $6, $7, $8
	) ON CONFLICT (name) DO UPDATE
	SET
		aliases = $2,
		description = $3,
		weight = $4,
		type = $5,
		ac = $6,
		worn = $7,
		attributes = $8
	RETURNING id
	;`

	var id int32
	err := r.pool.QueryRow(ctx, sql,
		armour.Name,
		armour.Aliases,
		armour.Description,
		armour.Weight,
		armour.Type,
		armour.AC,
		armour.Worn,
		armour.Attributes.AsMap()).Scan(&id)
	return id, err
}

func (r Repo) GetArmour(ctx context.Context, id int32) (models.Armour, error) {
	return r.getArmour(ctx, byIdClause, id)
}

func (r Repo) GetArmourByName(ctx context.Context, name string) (models.Armour, error) {
	return r.getArmour(ctx, byNameClause, name)
}

func (r Repo) getArmour(ctx context.Context, clause string, params ...any) (models.Armour, error) {
	const _sql = `SELECT
		id, name, aliases, description, weight, type, ac, worn, attributes
	FROM microdungeon.armour`

	var b strings.Builder
	b.WriteString(_sql)
	b.WriteString("\n")
	b.WriteString(clause)
	b.WriteString("\n;")

	sql := b.String()

	r.log.
		Debug("query", zap.String("sql", sql))

	var armour models.Armour

	err := r.pool.QueryRow(ctx, sql, params...).Scan(
		&armour.Id,
		&armour.Name,
		&armour.Aliases,
		&armour.Description,
		&armour.Weight,
		&armour.Type,
		&armour.AC,
		&armour.Worn,
		&armour.Attributes,
	)
	return armour, err
}
