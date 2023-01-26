package announcements

import (
	"encoding/json"
	"github.com/danilluk1/avito-tech/internal/app/api"
	"github.com/danilluk1/avito-tech/internal/app/api/api_errors"
	"github.com/danilluk1/avito-tech/internal/db/models"
	"github.com/danilluk1/avito-tech/internal/dto"
	"github.com/go-chi/chi/v5"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
)

//swagger:response announcementResponse
type AnnouncementResponse struct {
	//in: body
	Body struct {
		models.Announcement
	}
}

//swagger:response announcementsResponse
type AnnouncementsResponse struct {
	//in: body
	Body []struct {
		models.Announcement
	}
}

// swagger:route GET /api/announcements/{announcementId}
//
// # Get announcement by id
//
// Produces:
// - application/json
//
// Schemes: http, https
//
// Parameters:
//   - name: announcementId
//     in: path
//     description: id of announcement
//     required: true
//     type: integer
//
// Responses:
//
//	200: announcementResponse
//	404: notFoundError
//	500: internalError
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

// swagger:route GET /api/announcements
//
// # Get many announcements
//
// Produces:
//   - application/json
//
// Schemes: http, https
//
// Parameters:
//   - name: limit
//     in: query
//     description: limit
//     required: false
//     schema:
//     type: integer
//     max: 30
//     min: 1
//   - name: offset
//     in: query
//     description: offset
//     required: false
//     schema:
//     type: integer
//     min: 0
//   - name: sort
//     in: query
//     description: sort type
//     required: false
//     schema:
//     type: string
//   - name: order
//     in: query
//     description: order type
//     required: false
//     schema:
//     type: string
//
// Responses:
//
//	200: announcementsResponse
//	400: validationError
//	500: internalError
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
			switch dto.SortBy(sortByParam) {
			case dto.SortByDate:
				sortBy = dto.SortByDate
			case dto.SortByPrice:
				sortBy = dto.SortByPrice
			default:
				sortBy = ""
			}

			if sortBy == "" {
				response := api_errors.CreateBadRequestError([]string{"wrong sort param"})
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
		}

		orderByParam := query.Get("order")
		if len(orderByParam) != 0 {
			switch dto.OrderBy(orderByParam) {
			case dto.OrderByAsc:
				orderBy = dto.OrderByAsc
			case dto.OrderByDesc:
				orderBy = dto.OrderByDesc
			default:
				orderBy = ""
			}

			if orderBy == "" {
				response := api_errors.CreateBadRequestError([]string{"wrong order param"})
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
		}

		announcements, err := GetOrderedAnnouncements(app, uint64(limit), offset, sortBy, orderBy)
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(announcements)
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}

func PostAnnouncement(app *api.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := r.Context().Value("body").(*dto.CreateAnnouncement)
		if dto.Price.LessThanOrEqual(decimal.NewFromInt(0)) {
			response := api_errors.CreateBadRequestError([]string{"price must be higher or equals to 0"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
			return
		}

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
