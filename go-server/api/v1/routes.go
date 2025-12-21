package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/x022513319/asynchronous-processing-practice/api/queue"
	"github.com/x022513319/asynchronous-processing-practice/api/v1/tasks"
)

func Register(r chi.Router, publisher *queue.Publisher) {
	taskHandler := tasks.NewTaskHandler(publisher)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/tasks", taskHandler.Create)
		r.Get("/tasks/{id}", taskHandler.Get)
		r.Get("/tasks", taskHandler.List)
	})
}
