package internal

import (
	"ass3/pkg/store/postgres"
	"ass3/services/contact/internal/delivery"
	"ass3/services/contact/internal/repository"
	"ass3/services/contact/internal/use_case"
)

func NewContactRepository(db *postgres.Database) repository.ContactRepository {
	return &repository.ContactRepositoryImpl{db}
}

func NewContactUseCase(repo repository.ContactRepository) use_case.ContactUseCase {
	return use_case.ContactUseCaseImpl{repo}
}

func NewContactDelivery(useCase use_case.ContactUseCase) delivery.ContactDelivery {
	return delivery.NewContactDelivery(useCase)
}
