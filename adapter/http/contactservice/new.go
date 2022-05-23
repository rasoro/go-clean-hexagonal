package contactservice

import "github.com/rasoro/go-clean-hexagonal/core/domain"

type service struct {
	usecase domain.ContactUseCase
}

// New returns conttrac implementation of ContractService
func New(usecase domain.ContactUseCase) domain.ContactService {
	return &service{
		usecase: usecase,
	}
}
