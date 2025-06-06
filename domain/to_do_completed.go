package domain

import "github.com/google/uuid"

const ToDoCompletedEventType = "io.thenativeweb.to-do.completed"

type ToDoCompleted struct {
	ToDoID uuid.UUID `json:"-"`
}
