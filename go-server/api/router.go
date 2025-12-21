package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/x022513319/asynchronous-processing-practice/api/common"
	v1 "github.com/x022513319/asynchronous-processing-practice/api/v1"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// 共用
	r.Get("/health", common.Health)

	// versioned API
	r.Route("/api", func(r chi.Router) {
		v1.Register(r)
	})

	return r
}
