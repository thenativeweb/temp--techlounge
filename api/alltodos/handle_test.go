package alltodos_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/techlounge-to-do/api/alltodos"
)

func TestHandle(t *testing.T) {
	payload := &alltodos.ResponseBody{
		ToDos: []alltodos.ResponseBodyToDo{
			{ID: "1", Intention: "Buy milk", IsCompleted: false},
			{ID: "2", Intention: "Read book", IsCompleted: true},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rr := httptest.NewRecorder()

	handler := alltodos.Handle(payload)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var responseBody alltodos.ResponseBody
	err := json.NewDecoder(rr.Body).Decode(&responseBody)
	assert.NoError(t, err)

	assert.Len(t, responseBody.ToDos, 2)
	assert.Equal(
		t,
		alltodos.ResponseBodyToDo{ID: "1", Intention: "Buy milk", IsCompleted: false},
		responseBody.ToDos[0],
	)
	assert.Equal(
		t,
		alltodos.ResponseBodyToDo{ID: "2", Intention: "Read book", IsCompleted: true},
		responseBody.ToDos[1],
	)
}
