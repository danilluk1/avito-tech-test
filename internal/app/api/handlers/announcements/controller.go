package announcements

import (
	"encoding/json"
	"github.com/danilluk1/avito-tech/internal/app/api"
	"github.com/danilluk1/avito-tech/internal/app/api/api_errors"
	"github.com/danilluk1/avito-tech/internal/db/models"
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

		var announcement *models.Announcement
		if r.URL.Query().Get("fields") != "" {
			announcement, err = GetAnnouncementById(app, int32(announcementId), true)
		} else {
			announcement, err = GetAnnouncementById(app, int32(announcementId), false)
		}
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if announcement == nil {
			http.NotFound(w, r)
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

func GetAnnouncements(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var limit uint8 = 10
		var offset uint64 = 0
		var sortBy = dto.SortByDate
		var orderBy = dto.OrderByAsc

		query := r.URL.Query()
		limitParam := query.Get("limit")
		if len(limitParam) != 0 {
			newLimit, err := strconv.ParseUint(limitParam, 10, 64)
			if err != nil {
				response := api_errors.CreateBadRequestError([]string{"wrong limit"})
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
			if newLimit > 30 {
				response := api_errors.CreateBadRequestError([]string{"limit must be lower than 30"})
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
			limit = uint8(newLimit)
		}

		offsetParam := query.Get("offset")
		if len(offsetParam) != 0 {
			newOffset, err := strconv.ParseUint(offsetParam, 10, 64)
			if err != nil {
				response := api_errors.CreateBadRequestError([]string{"wrong offset"})
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}

			offset = newOffset
		}

		sortByParam := query.Get("sort")
		if len(sortByParam) != 0 {

		}
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
