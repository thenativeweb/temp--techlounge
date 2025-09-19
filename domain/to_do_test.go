package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/techlounge-to-do/domain"
)

func TestNewToDo(t *testing.T) {
	id := uuid.New()
	toDo := domain.NewToDo(id)

	assert.Equal(t, id, toDo.ID)
	assert.False(t, toDo.IsCompleted)
}

func TestMemorize(t *testing.T) {
	t.Run("returns a memorized event", func(t *testing.T) {
		id := uuid.New()
		toDo := domain.NewToDo(id)

		event, err := toDo.Memorize("Buy milk", "Remember to buy milk on the way home.")
		assert.NoError(t, err)
		assert.Equal(t, id, event.ToDoID)
		assert.Equal(t, "Buy milk", event.Intention)
		assert.Equal(t, "Remember to buy milk on the way home.", event.Details)
	})

	t.Run("returns an error if intention is empty", func(t *testing.T) {
		id := uuid.New()
		toDo := domain.NewToDo(id)

		_, err := toDo.Memorize("", "Some details")
		assert.Error(t, err)
	})
}
