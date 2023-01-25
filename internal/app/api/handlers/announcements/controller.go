package announcements

import (
	"encoding/json"
	"github.com/danilluk1/avito-tech/internal/app/api"
	"github.com/danilluk1/avito-tech/internal/app/api/api_errors"
	"github.com/danilluk1/avito-tech/internal/dto"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func GetAnnouncement(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		announcementId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			app.Logger.Error(err)
			response := api_errors.CreateBadRequestError([]string{"Id must be a number"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
			return
		}

		announcement, err := GetAnnouncementById(app, int32(announcementId))
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if announcement == nil {
			http.Error(w, "Can't find announcement with given id", http.StatusNotFound)
		}

		data, err := json.Marshal(announcement)
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}

func GetAnnouncements(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func PostAnnouncement(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := r.Context().Value("body").(*dto.CreateAnnouncement)
		if len(dto.Photos) > 3 {
			response := api_errors.CreateBadRequestError([]string{"You can't upload more than three photos"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
			return
		}
		announcement, err := AddAnnouncement(app, dto)
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(announcement)
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}
