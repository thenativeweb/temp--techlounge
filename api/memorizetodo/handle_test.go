package memorizetodo_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/eventsourcingdb-client-golang/eventsourcingdb"
	"github.com/thenativeweb/techlounge-to-do/api/memorizetodo"
	"github.com/thenativeweb/techlounge-to-do/domain"
)

func TestHandle(t *testing.T) {
	container := eventsourcingdb.NewContainer().
		WithImageTag("1.1.0").
		WithAPIToken("secret")

	err := container.Start(t.Context())
	assert.NoError(t, err)
	defer container.Stop(t.Context())

	client, err := container.GetClient(t.Context())
	assert.NoError(t, err)

	requestBody := `{"intention":"Buy milk","details":"Remember to buy milk on the way home."}`
	req := httptest.NewRequest(http.MethodPost, "/memorize-to-do", strings.NewReader(requestBody))
	rr := httptest.NewRecorder()

	handler := memorizetodo.Handle(client)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var responseBody memorizetodo.ResponseBody
	err = json.NewDecoder(rr.Body).Decode(&responseBody)
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, responseBody.ID)

	var events []eventsourcingdb.Event
	for event, err := range client.ReadEvents(
		t.Context(),
		"/",
		eventsourcingdb.ReadEventsOptions{Recursive: true},
	) {
		assert.NoError(t, err)
		events = append(events, event)
	}

	assert.Len(t, events, 1)

	event := events[0]
	assert.Equal(t, domain.ToDoMemorizedEventType, event.Type)

	var memorizedEvent domain.ToDoMemorized
	err = json.Unmarshal(event.Data, &memorizedEvent)
	assert.NoError(t, err)

	assert.Equal(t, "Buy milk", memorizedEvent.Intention)
	assert.Equal(t, "Remember to buy milk on the way home.", memorizedEvent.Details)
}
