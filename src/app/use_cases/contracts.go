package usecases

import (
	checkoutUc "fita/src/app/use_cases/checkout"
)

type AllUseCases struct {
	CheckoutUseCase checkoutUc.CheckoutUsecaseInterface
}
