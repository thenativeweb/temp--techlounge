package pendingtodos

import (
	"net/http"

	"github.com/thenativeweb/techlounge-to-do/httputil"
)

func Handle(pendingTodos *ResponseBody) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httputil.WriteJSONResponse(w, pendingTodos)
	}
}
