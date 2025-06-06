package alltodos

import (
	"net/http"

	"github.com/thenativeweb/techlounge-to-do/httputil"
)

func Handle(allTodos *ResponseBody) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httputil.WriteJSONResponse(w, allTodos)
	}
}
