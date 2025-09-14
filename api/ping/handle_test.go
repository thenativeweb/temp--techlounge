package ping_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/techlounge-to-do/api/ping"
)

func TestHandle(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rr := httptest.NewRecorder()

	handler := ping.Handle()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var responseBody ping.ResponseBody
	err := json.NewDecoder(rr.Body).Decode(&responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "pong", responseBody.Ping)
}
