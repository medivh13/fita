package checkout_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : fita
 */

import (
	"context"
	dtoCheckout "fita/src/app/dtos/items"
	"fita/src/domain/repositories"
	msHelper "fita/src/infra/helpers"
	"log"
)

type CheckoutUsecaseInterface interface {
	Checkout(ctx context.Context, checkoutDto *dtoCheckout.CheckoutReqDTO) (float64, error)
}

type chechkoutUseCase struct {
	ItemsRepo repositories.ItemsRepository
}

func NewCheckoutUseCase(itemsRepo repositories.ItemsRepository) *chechkoutUseCase {
	return &chechkoutUseCase{
		ItemsRepo: itemsRepo,
	}
}

func (uc *chechkoutUseCase) Checkout(ctx context.Context, checkoutDto *dtoCheckout.CheckoutReqDTO) (float64, error) {
	var sku []string

	for _, val := range checkoutDto.Data {
		sku = append(sku, val.Sku)
	}
	dataItems, err := uc.ItemsRepo.ListItems(ctx, sku)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	total, err := msHelper.CheckPromo(checkoutDto, dataItems)

	return total, nil
}
