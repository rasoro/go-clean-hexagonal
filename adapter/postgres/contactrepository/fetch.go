package contactrepository

import (
	"github.com/rasoro/go-clean-hexagonal/core/domain"
	"github.com/rasoro/go-clean-hexagonal/core/dto"
)

func (repository repository) Fetch(contactRequest *dto.FetchContactRequest) ([]domain.Contact, error) {
	// ctx := context.Background()
	contacts := []domain.Contact{}

	// err := repository.db.QueryRow(
	// 	ctx,
	// 	"SELECT * FROM contact",
	// ).Scan(&contacts)

	// if err != nil {
	// 	return nil, err
	// }

	// {
	// 	rows, err := repository.db.Query(
	// 		ctx,
	// 		"SELECT * FROM contact",
	// 	)
	// }

	return contacts, nil
}
