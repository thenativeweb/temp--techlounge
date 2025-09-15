package domain

import (
	"errors"

	"github.com/google/uuid"
	"github.com/thenativeweb/eventsourcingdb-client-golang/eventsourcingdb"
)

type ToDo struct {
	ID          uuid.UUID
	IsCompleted bool
}

func NewToDo(id uuid.UUID) *ToDo {
	return &ToDo{
		ID:          id,
		IsCompleted: false,
	}
}

func (t *ToDo) Memorize(intention string, details string) (*ToDoMemorized, error) {
	if len(intention) == 0 {
		return nil, errors.New("intention must not be empty")
	}
	if len(intention) > 255 {
		return nil, errors.New("intention must not be longer than 255 characters")
	}

	return &ToDoMemorized{
		ToDoID:    t.ID,
		Intention: intention,
		Details:   details,
	}, nil
}

func (t *ToDo) Complete() (*ToDoCompleted, error) {
	if t.IsCompleted {
		return nil, errors.New("to-do is already completed")
	}

	return &ToDoCompleted{
		ToDoID: t.ID,
	}, nil
}

// TODO: Don't hand over the database specific event struct,
// but rather a domain event struct which only contains the
// event's type and its payload.
func (t *ToDo) Apply(event eventsourcingdb.Event) {
	switch event.Type {
	case ToDoMemorizedEventType:
		// Intentionally left empty.
	case ToDoCompletedEventType:
		t.IsCompleted = true
	}
}
