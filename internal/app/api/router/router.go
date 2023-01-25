package api

import (
	"github.com/danilluk1/avito-tech/internal/app/api"
	"github.com/danilluk1/avito-tech/internal/app/api/handlers/announcements"
	"github.com/danilluk1/avito-tech/internal/app/api/middlewares"
	"github.com/danilluk1/avito-tech/internal/dto"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Setup(app *api.App) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api/announcement", func(r chi.Router) {
		r.Get("/", announcements.GetAnnouncements(app))
		r.Get("/{id}", announcements.GetAnnouncement(app))
		r.With(func(handler http.Handler) http.Handler {
			return middlewares.ValidateAndAttachBody(handler, app, &dto.CreateAnnouncement{})
		}).Post("/", announcements.PostAnnouncement(app))
	})

	return router
}
