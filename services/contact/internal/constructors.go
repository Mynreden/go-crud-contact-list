package internal

import (
	"ass3/services/contact/internal/delivery"
	"ass3/services/contact/internal/repository"
	"ass3/services/contact/internal/use_case"
)

func NewContactRepository() repository.ContactRepository {
	return nil
}

func NewContactUseCase(repo repository.ContactRepository) use_case.ContactUseCase {
	return nil
}

func NewContactDelivery(useCase use_case.ContactUseCase) delivery.ContactDelivery {
	return nil
}
