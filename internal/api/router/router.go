package api

import (
	"github.com/danilluk1/avito-tech/internal/api"
	"github.com/danilluk1/avito-tech/internal/api/handlers/announcements"
	"github.com/gorilla/mux"
)

func Setup(app *api.App) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/api/")
	r.HandleFunc("/announcement/{id}", announcements.GetAnnouncement(app)).Methods("GET")
	r.HandleFunc("/announcement", announcements.GetAnnouncements(app)).
		Methods("GET")
	r.HandleFunc("/announcement", announcements.PostAnnouncement(app)).Methods("POST")

	return r
}
