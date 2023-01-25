package api

import (
	"github.com/danilluk1/avito-tech/internal/api/handlers/announcements"
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/api/")
	r.HandleFunc("/announcement/{id}", announcements.GetAnnouncement).Methods("GET")
	r.HandleFunc("/announcement", announcements.GetAnnouncements).Methods("GET").Queries("count")
	r.HandleFunc("/announcement", announcements.PostAnnouncement).Methods("POST")

	return r
}
