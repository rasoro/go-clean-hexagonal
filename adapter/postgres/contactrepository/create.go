package contactrepository

import (
	"context"

	"github.com/rasoro/go-clean-hexagonal/core/domain"
	"github.com/rasoro/go-clean-hexagonal/core/dto"
)

func (repository repository) Create(contactRequest *dto.CreateContactRequest) (*domain.Contact, error) {
	ctx := context.Background()
	contact := domain.Contact{}

	err := repository.db.QueryRow(
		ctx,
		`INSERT INTO 
			contact (id, name, status, created_on, modified_on)
		VALUES
			(&1, &2, &3, NOW(), NOW())
		`,
		contactRequest.ID,
		contactRequest.Name,
		contactRequest.Status,
	).Scan(
		&contact.ID,
		&contact.Name,
		&contact.Status,
	)

	if err != nil {
		return nil, err
	}
	return &contact, nil
}
