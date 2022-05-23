package contactusecase

import (
	"github.com/rasoro/go-clean-hexagonal/core/domain"
	"github.com/rasoro/go-clean-hexagonal/core/dto"
)

func (u usecase) Fetch(contactRequest *dto.FetchContactRequest) ([]domain.Contact, error) {
	contacts, err := u.repository.Fetch(contactRequest)

	if err != nil {
		return nil, err
	}

	return contacts, nil
}
