package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : fita
 */

import (
	"context"

	repositories "fita/src/domain/repositories"
	models "fita/src/infra/models"

	"gorm.io/gorm"
)

type itemsRepository struct {
	connection *gorm.DB
}

func NewItemsRepository(db *gorm.DB) repositories.ItemsRepository {
	return &itemsRepository{
		connection: db,
	}
}

func (repo *itemsRepository) ListItems(ctx context.Context, sku []string) ([]*models.Items, error) {
	var itemsModel []*models.Items

	q := repo.connection.WithContext(ctx)
	if err := q.Raw(`SELECT id, sku, name, price,  inventory_qty FROM fita.items WHERE deleted_at IS NULL AND sku IN ?`, sku).Find(&itemsModel).Error; err != nil {
		return nil, err
	}

	return itemsModel, nil
}
