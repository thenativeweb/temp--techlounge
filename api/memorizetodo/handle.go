package memorizetodo

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/thenativeweb/eventsourcingdb-client-golang/eventsourcingdb"
	"github.com/thenativeweb/techlounge-to-do/domain"
	"github.com/thenativeweb/techlounge-to-do/httputil"
	"github.com/thenativeweb/techlounge-to-do/logging"
)

func Handle(client *eventsourcingdb.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBytes, err := io.ReadAll(r.Body)
		if err != nil {
			logging.Error("failed to read request body", "error", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			logging.Error("failed to unmarshal request body", "error", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id := uuid.New()
		intention := strings.TrimSpace(requestBody.Intention)
		details := strings.TrimSpace(requestBody.Details)

		// TODO: Everything that happens here should be done
		// in a separate layer that sits between the HTTP layer
		// and the domain layer. Loading the todo from the
		// database should happen there, calling the domain,
		// taking events back, and writing them to the database
		// should happen there. This is a bit of a mess right now.

		todo := domain.NewTodo(id)
		toDoMemorized, err := todo.Memorize(intention, details)
		if err != nil {
			logging.Error("failed to memorize todo", "error", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		subject := "/todo/" + toDoMemorized.ToDoID.String()
		_, err = client.WriteEvents([]eventsourcingdb.EventCandidate{
			{
				Source:  "https://to-do.thenativeweb.io",
				Subject: subject,
				Type:    domain.ToDoMemorizedEventType,
				Data:    toDoMemorized,
			},
		}, []eventsourcingdb.Precondition{
			eventsourcingdb.NewIsSubjectPristinePrecondition(subject),
		})

		// TODO: Beginning here, we have HTTP logic again, which
		// is fine. Basically, the above TODO ends here.

		if err != nil {
			logging.Error("failed to write event", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody := ResponseBody{
			ID: id,
		}
		httputil.WriteJSONResponse(w, responseBody)
	}
}
