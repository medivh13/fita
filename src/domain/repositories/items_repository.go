package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : fita
 */

import (
	"context"

	models "fita/src/infra/models"
)

type ItemsRepository interface {
	ListItems(context context.Context, sku []string) ([]*models.Items, error)
}
