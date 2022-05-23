package contactservice

import (
	"encoding/json"
	"net/http"

	"github.com/rasoro/go-clean-hexagonal/core/dto"
)

// Fetch returns list of contacts from a given request or default
func (s service) Fetch(w http.ResponseWriter, r *http.Request) {
	cr, err := dto.FromJSONFetchContactRequest(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contacts, err := s.usecase.Fetch(cr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(contacts)
}
