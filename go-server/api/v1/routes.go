package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/x022513319/asynchronous-processing-practice/api/v1/tasks"
)

func Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Post("/tasks", tasks.Create)
		r.Get("/tasks/{id}", tasks.Get)
		r.Get("/tasks", tasks.List)
	})
}
