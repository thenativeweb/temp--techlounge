package domain

import "github.com/google/uuid"

const ToDoMemorizedEventType = "io.thenativeweb.to-do.memorized"

type ToDoMemorized struct {
	ToDoID    uuid.UUID `json:"-"`
	Intention string    `json:"intention"`
	Details   string    `json:"details"`
}
