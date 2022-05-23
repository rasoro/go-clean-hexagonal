package domain

import (
	"net/http"
	"time"

	"github.com/rasoro/go-clean-hexagonal/core/dto"
)

// Contact is entity of table contact database
type Contact struct {
	ID         string
	Name       string
	Status     string
	URNs       []string
	CreatedOn  time.Time
	ModifiedOn time.Time
	LastSeenOn time.Time
}

// ContactService is a contract of http adapter layer
type ContactService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Fetch(w http.ResponseWriter, r *http.Request)
}

// ContactUseCase is a contract of business rule layer
type ContactUseCase interface {
	Create(contactRequest *dto.CreateContactRequest) (*Contact, error)
	Fetch(contactRequest *dto.FetchContactRequest) ([]Contact, error)
}

// ContactRepository is a contract of database connection adapter layer
type ContactRepository interface {
	Create(contactRequest *dto.CreateContactRequest) (*Contact, error)
	Fetch(contactRequest *dto.FetchContactRequest) ([]Contact, error)
}
