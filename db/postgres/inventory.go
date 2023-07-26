package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"justanother.org/microdungeon/internal/util"
	"justanother.org/microdungeon/models"
)

func (r Repo) storeInventory(ctx context.Context, inv []models.Object) ([]string, error) {
	items := make([]string, 0, len(inv))

	for _, _ob := range inv {
		switch ob := _ob.(type) {
		case models.Armour:
			id, err := r.storeArmour(ctx, ob)
			if err != nil {
				return nil, err
			}

			key := inventoryArmourPrefix + util.FormatID(id)
			items = append(items, key)
		case models.Weapon:
			return nil, errors.New("not yet implemented")
		default:
			return nil, errors.New("not yet implemented")
		}
	}

	return items, nil
}

func (r Repo) getInventory(ctx context.Context, items []string) ([]models.Object, error) {
	obs := make([]models.Object, 0, len(items))
	for _, str := range items {
		switch {
		case strings.HasPrefix(str, inventoryArmourPrefix):
			id, err := util.ParseID(strings.TrimPrefix(str, inventoryArmourPrefix))
			if err != nil {
				return nil, err
			}

			ob, err := r.GetArmour(ctx, id)
			if err != nil {
				return nil, err
			}

			obs = append(obs, ob)

		case strings.HasPrefix(str, inventoryWeaponPrefix):
			return nil, fmt.Errorf("%s, %w", str, errors.New("not yet implemented"))
		default:
			return nil, fmt.Errorf("%s, %w", str, errors.New("not yet implemented"))
		}
	}

	return obs, nil
}
