package contactrepository

import (
	"github.com/rasoro/go-clean-hexagonal/adapter/postgres"
	"github.com/rasoro/go-clean-hexagonal/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of contactRepository
func New(db postgres.PoolInterface) domain.ContactRepository {
	return &repository{
		db: db,
	}
}
