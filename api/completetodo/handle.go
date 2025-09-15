package completetodo

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/thenativeweb/eventsourcingdb-client-golang/eventsourcingdb"
	"github.com/thenativeweb/techlounge-to-do/domain"
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

		id := requestBody.ID

		// TODO: Everything that happens here should be done
		// in a separate layer that sits between the HTTP layer
		// and the domain layer. Loading the todo from the
		// database should happen there, calling the domain,
		// taking events back, and writing them to the database
		// should happen there. This is a bit of a mess right now.

		todo := domain.NewToDo(id)

		latestEventID := ""
		for event, err := range client.ReadEvents(
			r.Context(),
			"/todo/"+id.String(),
			eventsourcingdb.ReadEventsOptions{
				Recursive: false,
			},
		) {
			if err != nil {
				logging.Error("failed to read events", "error", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			todo.Apply(event)
			latestEventID = event.ID
		}

		toDoCompleted, err := todo.Complete()
		if err != nil {
			logging.Error("failed to complete todo", "error", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		subject := "/todo/" + toDoCompleted.ToDoID.String()
		_, err = client.WriteEvents([]eventsourcingdb.EventCandidate{
			{
				Source:  "https://to-do.thenativeweb.io",
				Subject: subject,
				Type:    domain.ToDoCompletedEventType,
				Data:    toDoCompleted,
			},
		}, []eventsourcingdb.Precondition{
			eventsourcingdb.NewIsSubjectOnEventIDPrecondition(subject, latestEventID),
		})

		// TODO: Beginning here, we have HTTP logic again, which
		// is fine. Basically, the above TODO ends here.

		if err != nil {
			logging.Error("failed to write event", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
