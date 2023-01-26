package announcements

import (
	"github.com/danilluk1/avito-tech/internal/app/api"
	"github.com/danilluk1/avito-tech/internal/db/models"
	"github.com/danilluk1/avito-tech/internal/dto"
)

func AddAnnouncement(app *api.App, dto *dto.CreateAnnouncement) (*models.Announcement, error) {
	announcement, err := app.AnnouncementService.Create(dto)
	if err != nil {
		return nil, err
	}

	return announcement, nil
}

func GetAnnouncementById(app *api.App, id int32, optional bool) (*models.Announcement, error) {
	announcement, err := app.AnnouncementService.GetById(id, optional)
	if err != nil {
		return nil, err
	}

	return announcement, nil
}

func GetOrderedAnnouncements(app *api.App, limit, offset uint64, sortBy dto.SortBy, orderBy dto.OrderBy) ([]models.Announcement, error) {
	announcements, err := app.AnnouncementService.GetMany(limit, offset, sortBy, orderBy)
	if err != nil {
		return nil, err
	}

	return announcements, nil
}
