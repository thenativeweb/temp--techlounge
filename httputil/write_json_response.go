package httputil

import (
	"encoding/json"
	"net/http"

	"github.com/thenativeweb/techlounge-to-do/logging"
)

func WriteJSONResponse(w http.ResponseWriter, responseBody any) {
	responseBytes, err := json.Marshal(responseBody)
	if err != nil {
		logging.Error("failed to marshal response body", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
}
