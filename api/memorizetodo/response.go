package memorizetodo

import "github.com/google/uuid"

type ResponseBody struct {
	ID uuid.UUID `json:"id"`
}
