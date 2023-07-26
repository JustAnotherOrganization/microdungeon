package postgres

import (
	"context"
	"strings"

	"go.uber.org/zap"
	"justanother.org/microdungeon/models"
)

func (r Repo) StoreClass(ctx context.Context, class models.Class) error {
	const sql = `
	INSERT INTO microdungeon.classes (
		name, hit_die, primary_abilities, saves
	) VALUES (
		$1, $2, $3, $4
	) ON CONFLICT (name) DO UPDATE
	SET
		hit_die = $2,
		primary_abilities = $3,
		saves = $4
	;`

	_, err := r.pool.Exec(ctx, sql,
		class.Name,
		class.HitDie,
		class.PrimaryAbilities,
		class.Saves)
	return err
}

func (r Repo) GetClass(ctx context.Context, id int32) (models.Class, error) {
	return r.getClass(ctx, byIdClause, id)
}

func (r Repo) GetClassByName(ctx context.Context, name string) (models.Class, error) {
	return r.getClass(ctx, byNameClause, name)
}

func (r Repo) getClass(ctx context.Context, clause string, params ...any) (models.Class, error) {
	const _sql = `SELECT id, name, hit_die, primary_abilities, saves FROM microdungeon.classes`

	var b strings.Builder
	b.WriteString(_sql)
	b.WriteString("\n")
	b.WriteString(clause)

	sql := b.String()

	r.log.
		Debug("query", zap.String("sql", sql))

	var class models.Class

	err := r.pool.QueryRow(ctx, sql, params...).Scan(
		&class.Id,
		&class.Name,
		&class.HitDie,
		&class.PrimaryAbilities,
		&class.Saves,
	)
	return class, err
}
