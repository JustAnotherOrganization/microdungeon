package postgres

import (
	"context"
	"strings"

	"go.uber.org/zap"
	"justanother.org/microdungeon/models"
)

const (
	inventoryArmourPrefix = "armour-"
	inventoryWeaponPrefix = "weapon-"
	inventoryObjectPrefix = "object-"
)

func (r Repo) StoreCharacter(ctx context.Context, char models.Character) error {
	abilitiesID, err := r.storeAbilities(ctx, char.Score.Abilities)
	if err != nil {
		return err
	}

	items, err := r.storeInventory(ctx, char.Inventory)
	if err != nil {
		return err
	}

	const sql = `
	INSERT INTO microdungeon.characters (
		name, player_name, race_name, class_name, level,
		alignment_ethical, alignment_moral, experience,
		abilities, attributes, inventory
	) VALUES (
		$1, $2, $3, $4, $5,
		$6, $7, $8,
		$9, $10, $11
	) ON CONFLICT (name) DO UPDATE
	SET
		player_name = $2,
		race_name = $3,
		class_name = $4,
		level = $5,
		alignment_ethical = $6,
		alignment_moral = $7,
		experience = $8,
		abilities = $9,
		attributes = $10,
		inventory = $11
	;`

	_, err = r.pool.Exec(ctx, sql,
		char.Name,
		char.PlayerName,
		char.Race.Name,
		char.Class.Name,
		char.Score.Level,
		char.Alignment.Ethical,
		char.Alignment.Moral,
		char.Score.Experience,
		abilitiesID,
		char.Attributes.AsMap(),
		items,
	)
	return err
}

func (r Repo) GetCharacter(ctx context.Context, id int32) (models.Character, error) {
	return r.getCharacter(ctx, byIdClause, id)
}

func (r Repo) GetCharacterByName(ctx context.Context, name string) (models.Character, error) {
	return r.getCharacter(ctx, byNameClause, name)
}

func (r Repo) getCharacter(ctx context.Context, clause string, params ...any) (models.Character, error) {
	const _sql = `
	SELECT
		id, name, player_name, race_name, class_name, level,
		alignment_ethical, alignment_moral, experience,
		abilities, attributes, inventory
	FROM microdungeon.characters`

	var b strings.Builder
	b.WriteString(_sql)
	b.WriteString("\n")
	b.WriteString(clause)
	b.WriteString("\n;")

	sql := b.String()

	r.log.
		Debug("query", zap.String("sql", sql))

	var (
		char        models.Character
		raceName    string
		className   string
		abilitiesID int32
		inv         []string
	)

	err := r.pool.QueryRow(ctx, sql, params...).Scan(
		&char.Id,
		&char.Name,
		&char.PlayerName,
		&raceName,
		&className,
		&char.Score.Level,
		&char.Alignment.Ethical,
		&char.Alignment.Moral,
		&char.Score.Experience,
		&abilitiesID,
		&char.Attributes,
		&inv,
	)
	if err != nil {
		return char, err
	}

	char.Race, err = r.GetRaceByName(ctx, raceName)
	if err != nil {
		return char, err
	}

	char.Class, err = r.GetClassByName(ctx, className)
	if err != nil {
		return char, err
	}

	char.Inventory, err = r.getInventory(ctx, inv)
	if err != nil {
		return char, err
	}

	char.Score.Abilities, err = r.getAbilities(ctx, abilitiesID)
	if err != nil {
		return char, err
	}

	return char, nil
}
