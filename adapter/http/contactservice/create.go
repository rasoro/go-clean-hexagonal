package contactservice

import (
	"encoding/json"
	"net/http"

	"github.com/rasoro/go-clean-hexagonal/core/dto"
)

// Create creates a new contact from a given request
func (s service) Create(w http.ResponseWriter, r *http.Request) {
	cr, err := dto.FromJSONCreateContactRequest(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c, err := s.usecase.Create(cr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(c)
}
