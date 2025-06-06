package ping

import (
	"net/http"

	"github.com/thenativeweb/techlounge-to-do/httputil"
)

func Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseBody := ResponseBody{
			Ping: "pong",
		}

		httputil.WriteJSONResponse(w, responseBody)
	}
}
