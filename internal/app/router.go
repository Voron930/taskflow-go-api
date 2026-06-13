package app

import (
	"net/http"

	"taskflow/internal/handler"
	"taskflow/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter(authHandler *handler.AuthHandler, taskHandler *handler.TaskHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", handler.HealthCheck)
	r.Post("/auth/register", authHandler.Register)
	r.Post("/auth/login", authHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Get("/profile", handler.Profile)
		r.Get("/tasks", taskHandler.GetTasks)
		r.Get("/tasks/{id}", taskHandler.GetTaskByID)
		r.Post("/tasks", taskHandler.CreateTask)
		r.Put("/tasks/{id}", taskHandler.UpdateTask)
		r.Delete("/tasks/{id}", taskHandler.DeleteTask)
	})

	return r
}
