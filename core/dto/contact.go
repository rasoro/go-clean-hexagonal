package dto

import (
	"encoding/json"
	"io"
)

type CreateContactRequest struct {
	ID     string   `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Status string   `json:"status,omitempty"`
	URNs   []string `json:"ur_ns,omitempty"`
}

type FetchContactRequest struct {
	ID     string   `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Status string   `json:"status,omitempty"`
	URNs   []string `json:"ur_ns,omitempty"`
}

// FromJSONCreateContactRequest converts json body request to a CreateContactRequest
func FromJSONCreateContactRequest(body io.Reader) (*CreateContactRequest, error) {
	createContactRequest := CreateContactRequest{}
	if err := json.NewDecoder(body).Decode(&createContactRequest); err != nil {
		return nil, err
	}

	return &createContactRequest, nil
}

// FromJSONFetchContactRequest converts json body request to a FetchContactRequest
func FromJSONFetchContactRequest(body io.Reader) (*FetchContactRequest, error) {
	fetchContactRequest := FetchContactRequest{}
	if err := json.NewDecoder(body).Decode(&fetchContactRequest); err != nil {
		return nil, err
	}

	return &fetchContactRequest, nil
}
