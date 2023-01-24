package api

import (
	"github.com/danilluk1/avito-tech/internal/api/handlers"
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/api/")
	r.HandleFunc("/announcement/{id}", handlers.GetAnnouncement).Methods("GET")
	r.HandleFunc("/announcement", handlers.GetAnnouncements).Queries("count")

	return &r
}
