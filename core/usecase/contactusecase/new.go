package contactusecase

import "github.com/rasoro/go-clean-hexagonal/core/domain"

type usecase struct {
	repository domain.ContactRepository
}

// New returns contract implementation of ContractUseCase
func New(repository domain.ContactRepository) domain.ContactUseCase {
	return &usecase{
		repository: repository,
	}
}
