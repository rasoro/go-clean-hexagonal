package di

import (
	"github.com/rasoro/go-clean-hexagonal/adapter/http/contactservice"
	"github.com/rasoro/go-clean-hexagonal/adapter/postgres"
	"github.com/rasoro/go-clean-hexagonal/adapter/postgres/contactrepository"
	"github.com/rasoro/go-clean-hexagonal/core/domain"
	"github.com/rasoro/go-clean-hexagonal/core/usecase/contactusecase"
)

func ConfigContactDI(conn postgres.PoolInterface) domain.ContactService {
	contactRepository := contactrepository.New(conn)
	contactUseCase := contactusecase.New(contactRepository)
	contactService := contactservice.New(contactUseCase)

	return contactService
}
