package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/thenativeweb/eventsourcingdb-client-golang/eventsourcingdb"
	"github.com/thenativeweb/techlounge-to-do/api/alltodos"
	"github.com/thenativeweb/techlounge-to-do/api/completetodo"
	"github.com/thenativeweb/techlounge-to-do/api/memorizetodo"
	"github.com/thenativeweb/techlounge-to-do/api/pendingtodos"
	"github.com/thenativeweb/techlounge-to-do/api/ping"
	"github.com/thenativeweb/techlounge-to-do/domain"
	"github.com/thenativeweb/techlounge-to-do/logging"
)

func main() {
	baseURLString := os.Getenv("ESDB_URL")
	if baseURLString == "" {
		baseURLString = "http://localhost:3000"
	}

	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		logging.Fatal("failed to parse base URL", "error", err)
	}

	apiToken := os.Getenv("ESDB_API_TOKEN")
	if apiToken == "" {
		apiToken = "secret"
	}

	client, err := eventsourcingdb.NewClient(baseURL, apiToken)
	if err != nil {
		logging.Fatal("failed to create eventsourcingdb client", "error", err)
	}

	logging.Info("pinging eventsourcingdb...")
	err = client.Ping()
	if err != nil {
		logging.Fatal("failed to ping eventsourcingdb", "error", err)
	}
	logging.Info("eventsourcingdb is up and running")

	mux := http.NewServeMux()

	// Status
	mux.HandleFunc(
		"GET /api/v1/ping",
		ping.Handle(),
	)

	// Commands
	mux.HandleFunc(
		"POST /api/v1/memorize-to-do",
		memorizetodo.Handle(client),
	)

	mux.HandleFunc(
		"POST /api/v1/complete-to-do",
		completetodo.Handle(client),
	)

	// Queries
	allTodos := &alltodos.ResponseBody{
		ToDos: []alltodos.ResponseBodyToDo{},
	}

	pendingTodos := &pendingtodos.ResponseBody{
		ToDos: []pendingtodos.ResponseBodyToDo{},
	}

	// TODO: This should be located somewhere else, it's here
	// only for the sake of simplicity.

	go func() {
		for event, err := range client.ObserveEvents(
			context.Background(),
			"/",
			eventsourcingdb.ObserveEventsOptions{
				Recursive: true,
			},
		) {
			if err != nil {
				logging.Fatal("failed to observe events", "error", err)
			}

			switch event.Type {
			case domain.ToDoMemorizedEventType:
				var memorizedEvent domain.ToDoMemorized
				err := json.Unmarshal(event.Data, &memorizedEvent)
				if err != nil {
					logging.Error("failed to unmarshal event", "error", err)
					continue
				}

				allTodos.ToDos = append(allTodos.ToDos, alltodos.ResponseBodyToDo{
					ID:          path.Base(event.Subject),
					Intention:   memorizedEvent.Intention,
					IsCompleted: false,
				})

				pendingTodos.ToDos = append(pendingTodos.ToDos, pendingtodos.ResponseBodyToDo{
					ID:        path.Base(event.Subject),
					Intention: memorizedEvent.Intention,
					Details:   memorizedEvent.Details,
				})

			case domain.ToDoCompletedEventType:
				id := path.Base(event.Subject)

				for i, todo := range allTodos.ToDos {
					if todo.ID == id {
						allTodos.ToDos[i].IsCompleted = true
						break
					}
				}

				for i, todo := range pendingTodos.ToDos {
					if todo.ID == id {
						pendingTodos.ToDos = append(pendingTodos.ToDos[:i], pendingTodos.ToDos[i+1:]...)
						break
					}
				}
			}
		}
	}()

	mux.HandleFunc(
		"GET /api/v1/all-to-dos",
		alltodos.Handle(allTodos),
	)

	mux.HandleFunc(
		"GET /api/v1/pending-to-dos",
		pendingtodos.Handle(pendingTodos),
	)

	// UI

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	addr := fmt.Sprintf(":%s", port)

	logging.Info("server starting...", "port", port)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		logging.Fatal("failed to start server", "error", err)
	}
}
