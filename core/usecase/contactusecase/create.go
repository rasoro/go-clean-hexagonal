package contactusecase

import (
	"github.com/rasoro/go-clean-hexagonal/core/domain"
	"github.com/rasoro/go-clean-hexagonal/core/dto"
)

func (usecase usecase) Create(contactRequest *dto.CreateContactRequest) (*domain.Contact, error) {
	contact, err := usecase.repository.Create(contactRequest)

	if err != nil {
		return nil, err
	}

	return contact, nil
}
