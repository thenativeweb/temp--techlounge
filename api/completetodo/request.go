package completetodo

import "github.com/google/uuid"

type RequestBody struct {
	ID uuid.UUID `json:"id"`
}
